package event

import (
	"reflect"
	"sync"
)

type Store map[uint8][]interface{}

type Dispatcher struct {
	sync.RWMutex

	events Store
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		events: make(Store),
	}
}

func (e *Dispatcher) On(eventId uint8, fn interface{}) {
	e.Lock()
	defer e.Unlock()

	e.events[eventId] = append(e.events[eventId], fn)
}

func (e *Dispatcher) Dispatch(eventId uint8, packet interface{}) {
	e.RLock()
	defer e.RUnlock()

	fns := e.events[eventId]
	for _, fn := range fns {
		e.call(fn, packet)
	}
}

func (e *Dispatcher) call(fn interface{}, packet interface{}) {
	// no parameter or type control, direct call
	reflect.ValueOf(fn).Call([]reflect.Value{
		reflect.ValueOf(packet),
	})
}
