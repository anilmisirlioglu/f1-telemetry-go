package math

import "testing"

func TestMath_Sqrt32(t *testing.T) {
	tests := []struct {
		n, want float32
	}{
		{4, 2},
		{9, 3},
		{16, 4},
		{100, 10},
	}
	for _, c := range tests {
		sqrt := sqrt32(c.n)
		if sqrt != c.want {
			t.Errorf("got: %g, want: %g", sqrt, c.want)
		}
	}
}

func TestMath_Pow32(t *testing.T) {
	tests := []struct {
		n1, n2 float32
		y      float64
		want   float32
	}{
		{8, 4, 2, 16},
		{200, 100, 2, 10000},
		{5, 3, 3, 8},
	}
	for _, c := range tests {
		pow := pow32(c.n1, c.n2, c.y)
		if pow != c.want {
			t.Errorf("got: %g, want: %g", pow, c.want)
		}
	}
}
