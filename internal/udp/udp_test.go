package udp

import (
	"bytes"
	"encoding/binary"
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/env"
	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/env/event"
	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/packets"
)

type EventTestCase struct {
	code       string
	packetType reflect.Type // expected packet type
}

type PacketTestCase struct {
	id         uint8
	packetType reflect.Type
}

var (
	testAddr = &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 20777,
	}
	testTimeout = time.Second
)

func NewMockServerAndClient(t *testing.T) (*Server, net.Conn, func()) {
	serv, err := Serve(testAddr)
	if err != nil {
		t.Fatal(err)
	}

	conn, err := net.DialTimeout("udp", testAddr.String(), testTimeout)
	if err != nil {
		t.Fatal("could not connect to udp server: ", err)
	}

	tOut := time.Now().Add(testTimeout)
	_ = serv.conn.SetReadDeadline(tOut)
	_ = conn.SetWriteDeadline(tOut)

	return serv, conn, func() {
		defer serv.conn.Close()
		defer conn.Close()
	}
}

func TestUDP_Serve(t *testing.T) {
	serv, conn, f := NewMockServerAndClient(t)
	defer f()

	wBytes := []byte("UDP TEST")
	if _, err := conn.Write(wBytes); err != nil {
		t.Fatal(err)
	}

	rBytes := make([]byte, len(wBytes))
	if _, err := serv.conn.Read(rBytes); err != nil {
		t.Fatal(err)
	}
}

func TestUDP_ReadSocket(t *testing.T) {
	var packetMap = map[uint8]interface{}{
		env.PacketMotion: &packets.PacketMotionData{
			Header: packets.PacketHeader{
				PacketFormat:            2020,
				GameMajorVersion:        1,
				GameMinorVersion:        18,
				PacketVersion:           1,
				PacketID:                0,
				SessionUID:              11855624319420004949,
				SessionTime:             1.5710948,
				FrameIdentifier:         43,
				PlayerCarIndex:          19,
				SecondaryPlayerCarIndex: 255,
			},
			CarMotionData:          [22]packets.CarMotionData{}, // empty
			SuspensionPosition:     [4]float32{-0.6432814, 0.049691506, -0.122375205, -0.11044062},
			SuspensionVelocity:     [4]float32{-5.7109776, -2.5368745, -0.33160102, 1.6033937},
			SuspensionAcceleration: [4]float32{-742.09515, -298.6635, -38.07375, 275.05835},
			WheelSpeed:             [4]float32{120.0, 120.0, 120.0, 120.0},
			WheelSlip:              [4]float32{10.5, 5.3, 9.5, 9.5},
			LocalVelocityX:         -0.00016372708,
			LocalVelocityY:         -0.0011803857,
			LocalVelocityZ:         0.0015708461,
			AngularVelocityX:       0.0113855535,
			AngularVelocityY:       -0.00053515914,
			AngularVelocityZ:       0.0073023424,
			AngularAccelerationX:   1.5726409,
			AngularAccelerationY:   0.047205403,
			AngularAccelerationZ:   1.0070984,
			FrontWheelsAngle:       0,
		},
	}

	serv, conn, f := NewMockServerAndClient(t)
	defer f()

	for id, p := range packetMap {
		var err error
		buf := new(bytes.Buffer)
		err = binary.Write(buf, binary.LittleEndian, p)
		_, err = conn.Write(buf.Bytes())
		if err != nil {
			t.Error("packet write error: ", err)
			continue
		}

		h, want, err := serv.ReadSocket()
		if err != nil {
			t.Error("socket read error: ", err)
			continue
		}

		if h.PacketID != id {
			t.Errorf("got: %d, want: %d", h.PacketID, id)
			continue
		}

		if !reflect.DeepEqual(p, want) {
			t.Errorf("got: %v, want: %v", p, want)
		}
	}
}

func NewEventTestCase(code string, i interface{}) EventTestCase {
	return EventTestCase{code, reflect.TypeOf(i)}
}

func TestUDP_ResolveEventDetails(t *testing.T) {
	var events = []EventTestCase{
		NewEventTestCase(event.SessionStarted, nil),
		NewEventTestCase(event.SessionEnded, nil),
		NewEventTestCase(event.FastestLap, new(packets.FastestLap)),
		NewEventTestCase(event.Retirement, new(packets.Retirement)),
		NewEventTestCase(event.DRSEnabled, nil),
		NewEventTestCase(event.DRSDisabled, nil),
		NewEventTestCase(event.TeamMateInPit, new(packets.TeamMateInPits)),
		NewEventTestCase(event.ChequeredFlag, nil),
		NewEventTestCase(event.RaceWinner, new(packets.RaceWinner)),
		NewEventTestCase(event.PenaltyIssued, new(packets.Penalty)),
		NewEventTestCase(event.SpeedTrapTriggered, new(packets.SpeedTrap)),
	}

	for _, e := range events {
		var code [4]uint8
		copy(code[:], e.code)
		details := resolveEventDetails(&packets.PrePacketEventData{
			EventStringCode: code,
		})

		got := reflect.TypeOf(details)
		if got != e.packetType {
			t.Errorf("got: %+v, want: %+v", got, e.packetType)
		}
	}
}

func NewPacketTestCase(id uint8, i interface{}) PacketTestCase {
	return PacketTestCase{id, reflect.TypeOf(i)}
}

func TestUDP_NewPacketById(t *testing.T) {
	var packs = []PacketTestCase{
		NewPacketTestCase(env.PacketMotion, new(packets.PacketMotionData)),
		NewPacketTestCase(env.PacketSession, new(packets.PacketSessionData)),
		NewPacketTestCase(env.PacketLap, new(packets.PacketLapData)),
		NewPacketTestCase(env.PacketEvent, new(packets.PrePacketEventData)),
		NewPacketTestCase(env.PacketParticipants, new(packets.PacketParticipantsData)),
		NewPacketTestCase(env.PacketCarStatus, new(packets.PacketCarStatusData)),
		NewPacketTestCase(env.PacketCarTelemetry, new(packets.PacketCarTelemetryData)),
		NewPacketTestCase(env.PacketCarStatus, new(packets.PacketCarStatusData)),
		NewPacketTestCase(env.PacketFinalClassification, new(packets.PacketFinalClassificationData)),
		NewPacketTestCase(env.PacketLobbyInfo, new(packets.PacketLobbyInfoData)),
	}

	for _, p := range packs {
		pack := newPacketById(p.id)

		got := reflect.TypeOf(pack)
		if got != p.packetType {
			t.Errorf("got: %+v, want: %+v", got, p.packetType)
		}
	}
}
