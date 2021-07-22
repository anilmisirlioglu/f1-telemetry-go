package event

import (
	"context"
	"math"
	"reflect"
	"testing"
	"time"
)

const (
	TestEventA = iota
	TestEventB
	TestEventC
	TestEventD
)

func TestDispatcher_On(t *testing.T) {
	dispatcher := NewDispatcher()
	eventMap := map[uint8]int{
		TestEventA: 8,
		TestEventB: 1,
		TestEventC: 3,
		TestEventD: 1,
	}

	for eventId, repeatCount := range eventMap {
		for i := 0; i < repeatCount; i++ {
			dispatcher.On(eventId, func() {})
		}
	}

	for eventId, fns := range dispatcher.events {
		want := eventMap[eventId]
		got := len(fns)
		if want != got {
			t.Errorf("eventId=%d, got: %d, want: %d", eventId, got, want)
		}
	}
}

func TestDispatcher_Dispatch(t *testing.T) {
	type TestPacket struct {
		eventId uint8
		name    string
		payload interface{}
	}

	dispatcher := NewDispatcher()
	eventMap := map[uint8]*TestPacket{
		TestEventA: {TestEventA, "test-event-a", "string payload"},
		TestEventB: {TestEventB, "test-event-b", math.MaxFloat64},
		TestEventC: {TestEventC, "test-event-c", []byte("byte payload")},
		TestEventD: {TestEventD, "test-event-d", struct {
			str string
			int int
			arr []interface{}
		}{"foo bar", math.MinInt32, []interface{}{"lorem", "ipsum", 1, 2, 3}}},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	i := 4
	for eventId, packet := range eventMap {
		dispatcher.On(eventId, func(p *TestPacket) {
			if !reflect.DeepEqual(packet, p) {
				t.Errorf("event=%d, got: %+v, want: %+v", eventId, p, packet)
			} else {
				i--
			}
		})
		dispatcher.Dispatch(eventId, packet)
	}

	for {
		select {
		case <-ctx.Done():
			t.Fatalf("test failed with timeout")
		default:
			if i == 0 {
				return // test success
			}
		}
	}
}
