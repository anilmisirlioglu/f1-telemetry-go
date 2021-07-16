package telemetry

import (
	"log"
	"net"

	"github.com/anilmisirlioglu/f1-telemetry-go/internal/event"
	"github.com/anilmisirlioglu/f1-telemetry-go/internal/udp"
	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/env"
	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/packets"
)

type Client struct {
	server     *udp.Server
	Stats      *udp.Stats
	dispatcher *event.Dispatcher
}

const localhost = "127.0.0.1"

func NewClient() (*Client, error) {
	// default F1 game UDP port is 20777
	return NewClientByCustomIpAddressAndPort(localhost, 20777)
}

func NewClientByCustomPort(port int) (*Client, error) {
	return NewClientByCustomIpAddressAndPort(localhost, port)
}

func NewClientByCustomIpAddressAndPort(ipAddress string, port int) (*Client, error) {
	serv, err := udp.ServeUDP(&net.UDPAddr{
		IP:   net.ParseIP(ipAddress),
		Port: port,
	})
	if err != nil {
		return nil, err
	}

	client := &Client{
		Stats:      udp.NewStats(),
		server:     serv,
		dispatcher: event.NewDispatcher(),
	}
	return client, nil
}

func (c *Client) Run() {
	for {
		h, p, err := c.server.ReadSocket()
		if err != nil {
			c.Stats.IncErr()
			log.Println(err)
		}

		c.Stats.IncRecv()
		c.dispatcher.Dispatch(h.PacketID, p)
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
