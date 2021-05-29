package internal

import (
	"log"
	"net"

	"github.com/anilmisirlioglu/f1-telemetry/internal/env"
	"github.com/anilmisirlioglu/f1-telemetry/internal/event"
	"github.com/anilmisirlioglu/f1-telemetry/internal/packets"
)

func ServeUDP(addr *net.UDPAddr) (*net.UDPConn, error) {
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func ListenUDPSocket(dispatcher *event.Dispatcher, conn *net.UDPConn) {
	buf := make([]byte, 1024*5)
	_, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		log.Printf("read error: %s\n", err)
		return
	}

	header := new(packets.PacketHeader)
	if err = ReadPacket(&buf, header); err != nil {
		log.Println(err)
		return
	}

	pack := newPacketById(header.PacketID)
	if pack == nil {
		log.Printf("invalid packet: %d\n", header.PacketID)
		return
	}

	if err = ReadPacket(&buf, pack); err != nil {
		log.Printf("%d: %s\n", header.PacketID, err)
		return
	}

	dispatcher.Dispatch(header.PacketID, pack)
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
