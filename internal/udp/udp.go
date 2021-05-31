package udp

import (
	"errors"
	"fmt"
	"net"

	"github.com/anilmisirlioglu/f1-telemetry/internal"
	"github.com/anilmisirlioglu/f1-telemetry/internal/env"
	"github.com/anilmisirlioglu/f1-telemetry/internal/packets"
)

type Server struct {
	conn *net.UDPConn
}

func ServeUDP(addr *net.UDPAddr) (*Server, error) {
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}

	s := &Server{conn: conn}
	return s, nil
}

func (s *Server) ReadSocket() (*packets.PacketHeader, interface{}, error) {
	buf := make([]byte, 1024*5)
	_, _, err := s.conn.ReadFromUDP(buf)
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("read error: %s\n", err))
	}

	header := new(packets.PacketHeader)
	if err = internal.ReadPacket(buf, header); err != nil {
		return nil, nil, err
	}

	pack := newPacketById(header.PacketID)
	if pack == nil {
		return nil, nil, errors.New(fmt.Sprintf("invalid packet: %d\n", header.PacketID))
	}

	if err = internal.ReadPacket(buf, pack); err != nil {
		return nil, nil, errors.New(fmt.Sprintf("%d: %s\n", header.PacketID, err))
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
		return new(packets.PacketEventData)
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
