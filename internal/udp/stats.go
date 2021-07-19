package udp

import (
	"sync"
	"time"
)

type Stats struct {
	sync.Mutex

	t    time.Time
	recv int // Success UDP Packet Count
	err  int // Fail UDP Packet Count
	pps  int // Packet Per Second
}

func NewStats() *Stats {
	return &Stats{
		recv: 0,
		err:  0,
		pps:  0,
		t:    time.Now(),
	}
}

func (s *Stats) RecvCount() int {
	return s.recv
}

func (s *Stats) ErrCount() int {
	return s.err
}

// PPS = Packet Per Second
func (s *Stats) PPS() int {
	s.Lock()
	defer s.Unlock()
	if s.eqPPSTime() {
		s.t = time.Now()
		s.pps = 0
	}

	return s.pps
}

func (s *Stats) IncRecv() {
	s.Lock()
	s.recv += 1
	s.Unlock()

	s.resolvePPS()
}

func (s *Stats) IncErr() {
	s.Lock()
	s.err += 1
	s.Unlock()
}

func (s *Stats) resolvePPS() {
	s.Lock()
	defer s.Unlock()
	if s.eqPPSTime() {
		s.t = time.Now()
		s.pps = 0
	} else {
		s.pps += 1
	}
}

func (s *Stats) eqPPSTime() bool {
	now := time.Now()
	return s.t.Second() != now.Second() || s.t.Second() == now.Second() && s.t.After(now.Add(time.Second))
}
