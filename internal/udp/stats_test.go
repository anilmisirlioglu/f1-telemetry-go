package udp

import (
	"testing"
	"time"
)

func TestStats_Inc(t *testing.T) {
	s := NewStats()
	t.Run("Recv", func(t *testing.T) {
		want := 50000
		for i := 0; i < want; i++ {
			s.IncRecv()
		}

		got := s.RecvCount()
		if want != got {
			t.Errorf("want: %d, got: %d", want, got)
		}
	})
	t.Run("Err", func(t *testing.T) {
		want := 15000
		for i := 0; i < want; i++ {
			s.IncErr()
		}

		got := s.ErrCount()
		if want != got {
			t.Errorf("want: %d, got: %d", want, got)
		}
	})
}

func TestStats_eqPPSTime(t *testing.T) {
	s := NewStats()
	if s.eqPPSTime() {
		t.Fatalf("PPS should be equal")
	}

	time.Sleep(time.Second)

	if !s.eqPPSTime() {
		t.Fatalf("PPS should not be equal")
	}
}

func TestStats_PPS(t *testing.T) {
	time.Sleep(time.Second - time.Duration(time.Now().Nanosecond()))

	s := NewStats()
	got := s.PPS()

	ticker := time.NewTicker(time.Nanosecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				break
			case <-ticker.C:
				s.IncRecv()
				got = s.PPS()
			}
		}
	}()

	time.Sleep(time.Millisecond * 500)
	ticker.Stop()
	done <- true

	want := s.RecvCount()
	if want != got {
		t.Fatalf("want: %d, got: %d", want, got)
	}

	time.Sleep(time.Millisecond * 500)
	got = s.PPS()
	if s.PPS() != 0 {
		t.Fatalf("want: 0, got: %d", got)
	}
}
