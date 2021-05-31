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
	ppr  int // Packet Per Second
}

func NewStats() *Stats {
	return &Stats{
		recv: 0,
		err:  0,
		ppr:  0,
		t:    time.Now(),
	}
}

func (s *Stats) RecvCount() int {
	return s.recv
}

func (s *Stats) ErrCount() int {
	return s.err
}

// PPR = Packet Per Second
func (s *Stats) PPR() int {
	s.Lock()
	defer s.Unlock()
	if s.eqPPRTime() {
		s.t = time.Now()
		s.ppr = 0
	}

	return s.ppr
}

func (s *Stats) IncRecv() {
	s.Lock()
	s.recv += 1
	s.Unlock()

	s.resolvePPR()
}

func (s *Stats) IncErr() {
	s.Lock()
	s.err += 1
	s.Unlock()
}

func (s *Stats) resolvePPR() {
	s.Lock()
	defer s.Unlock()
	if s.eqPPRTime() {
		s.t = time.Now()
		s.ppr = 0
	} else {
		s.ppr += 1
	}
}

func (s *Stats) eqPPRTime() bool {
	now := time.Now()
	return s.t.Second() != now.Second() || s.t.Second() == now.Second() && s.t.After(now.Add(time.Second))
}
