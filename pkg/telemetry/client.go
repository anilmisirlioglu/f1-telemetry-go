package telemetry

import (
	"net"

	"github.com/anilmisirlioglu/f1-telemetry/internal"
	"github.com/anilmisirlioglu/f1-telemetry/internal/env"
	"github.com/anilmisirlioglu/f1-telemetry/internal/event"
	"github.com/anilmisirlioglu/f1-telemetry/internal/packets"
)

type Client struct {
	conn       *net.UDPConn
	dispatcher *event.Dispatcher
}

func NewClient() (*Client, error) {
	port := 20777 // default F1 2020 UDP port
	return NewClientByCustomPort(&port)
}

func NewClientByCustomPort(port *int) (*Client, error) {
	conn, err := internal.ServeUDP(&net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: *port,
	})
	if err != nil {
		return nil, err
	}

	client := &Client{
		conn:       conn,
		dispatcher: event.NewEventDispatcher(),
	}
	return client, nil
}

func (c *Client) Run() {
	for {
		internal.ListenUDPSocket(c.dispatcher, c.conn)
	}
}

// Shared Events

func (c *Client) OnMotionPacket(fn func(packet *packets.PacketMotionData)) {
	c.dispatcher.On(env.PacketMotion, fn)
}

func (c *Client) OnSessionPacket(fn func(packet *packets.PacketSessionData)) {
	c.dispatcher.On(env.PacketSession, fn)
}

func (c *Client) OnLapPacket(fn func(packet *packets.PacketLapData)) {
	c.dispatcher.On(env.PacketLap, fn)
}

func (c *Client) OnEventPacket(fn func(packet *packets.PacketEventData)) {
	c.dispatcher.On(env.PacketEvent, fn)
}

func (c *Client) OnParticipantsPacket(fn func(packet *packets.PacketParticipantsData)) {
	c.dispatcher.On(env.PacketParticipants, fn)
}

func (c *Client) OnCarSetupPacket(fn func(packet *packets.PacketCarSetupData)) {
	c.dispatcher.On(env.PacketCarSetup, fn)
}

func (c *Client) OnCarTelemetryPacket(fn func(packet *packets.PacketCarTelemetryData)) {
	c.dispatcher.On(env.PacketCarTelemetry, fn)
}

func (c *Client) OnCarStatusPacket(fn func(packet *packets.PacketCarStatusData)) {
	c.dispatcher.On(env.PacketCarStatus, fn)
}

func (c *Client) OnFinalClassificationPacket(fn func(packet *packets.PacketFinalClassificationData)) {
	c.dispatcher.On(env.PacketFinalClassification, fn)
}

func (c *Client) OnLobbyInfoPacket(fn func(packet *packets.PacketLobbyInfoData)) {
	c.dispatcher.On(env.PacketLobbyInfo, fn)
}
