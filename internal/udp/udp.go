package udp

import (
	"fmt"
	"net"
	"unsafe"

	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/env"
	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/env/event"
	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/packets"
)

type Server struct {
	conn *net.UDPConn
}

func Serve(addr *net.UDPAddr) (*Server, error) {
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}

	s := &Server{conn: conn}
	return s, nil
}

func (s *Server) ReadSocket() (*packets.PacketHeader, interface{}, error) {
	buf := make([]byte, 1024+1024/2)
	_, _, err := s.conn.ReadFromUDP(buf)
	if err != nil {
		return nil, nil, fmt.Errorf(fmt.Sprintf("read error: %s", err))
	}

	header := new(packets.PacketHeader)
	if err = ReadPacket(buf, header); err != nil {
		return nil, nil, err
	}

	pack := newPacketById(header.PacketID)
	if pack == nil {
		return nil, nil, fmt.Errorf(fmt.Sprintf("invalid packet: %d", header.PacketID))
	}

	if err = ReadPacket(buf, pack); err != nil {
		return nil, nil, fmt.Errorf(fmt.Sprintf("%d: %s", header.PacketID, err))
	}

	if header.PacketID == env.PacketEvent {
		details := resolveEventDetails(pack.(*packets.PrePacketEventData))
		pre := pack.(*packets.PrePacketEventData)
		if details != nil {
			err = ReadPacket(pre.EventDetails[:unsafe.Sizeof(details)], details)
			if err != nil {
				return nil, nil, fmt.Errorf(fmt.Sprintf("event packet details read error: %s", err))
			}
		}
		pack = &packets.PacketEventData{
			Header:          pre.Header,
			EventStringCode: pre.EventStringCode,
			EventDetails:    details,
		}
	}

	return header, pack, nil
}

func newPacketById(packetId uint8) interface{} {
	switch packetId {
	case env.PacketMotion:
		return new(packets.PacketMotionData)
	case env.PacketSession:
		return new(packets.PacketSessionData)
	case env.PacketLap:
		return new(packets.PacketLapData)
	case env.PacketEvent:
		return new(packets.PrePacketEventData)
	case env.PacketParticipants:
		return new(packets.PacketParticipantsData)
	case env.PacketCarSetup:
		return new(packets.PacketCarSetupData)
	case env.PacketCarTelemetry:
		return new(packets.PacketCarTelemetryData)
	case env.PacketCarStatus:
		return new(packets.PacketCarStatusData)
	case env.PacketFinalClassification:
		return new(packets.PacketFinalClassificationData)
	case env.PacketLobbyInfo:
		return new(packets.PacketLobbyInfoData)
	}

	return nil
}

func resolveEventDetails(pre *packets.PrePacketEventData) interface{} {
	switch string(pre.EventStringCode[:]) {
	case event.FastestLap:
		return new(packets.FastestLap)
	case event.Retirement:
		return new(packets.Retirement)
	case event.TeamMateInPit:
		return new(packets.TeamMateInPits)
	case event.RaceWinner:
		return new(packets.RaceWinner)
	case event.PenaltyIssued:
		return new(packets.Penalty)
	case event.SpeedTrapTriggered:
		return new(packets.SpeedTrap)
	}

	return nil
}
