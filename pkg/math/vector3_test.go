package math

import (
	"math"
	"testing"
)

func TestVector3_Equals(t *testing.T) {
	tests := []*Vector3{
		NewVector3(0, 0, 0),
		NewVector3(math.MaxFloat32, math.MaxFloat32, math.MaxFloat32),
		NewVector3(-1, -2, -3),
		NewVector3(10, -10, 0),
		NewVector3(1e-12, 1e-14, 1e-16),
	}
	for _, test := range tests {
		if !test.Equals(test) {
			t.Errorf("%+v should be equal to itself", test)
		}
	}
}

func TestVector3_Add(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		tests := []struct {
			v       *Vector3
			x, y, z float32
			want    *Vector3
		}{
			{NewVector3(0, 0, 0), 1, 2, 3, NewVector3(1, 2, 3)},
			{NewVector3(1, 1, 1), 9, 9, 9, NewVector3(10, 10, 10)},
			{NewVector3(16, 16, 16), -6, -11, -13, NewVector3(10, 5, 3)},
			{NewVector3(-3, -5, 3), 3, 5, -3, NewVector3(0, 0, 0)},
			{NewVector3(1e-12, 1e-14, 1e-16), 0, 0, 0, NewVector3(1e-12, 1e-14, 1e-16)},
		}
		t.Run("from XYZ", func(t *testing.T) {
			for _, test := range tests {
				vec := test.v.Add(test.x, test.y, test.z)
				if !vec.Equals(test.want) {
					t.Errorf("v: %+v, got: %+v, want: %+v", test.v, vec, test.want)
				}
			}
		})
		t.Run("from Vector", func(t *testing.T) {
			for _, test := range tests {
				vec := test.v.AddFromVector3(NewVector3(test.x, test.y, test.z))
				if !vec.Equals(test.want) {
					t.Errorf("v: %+v, got: %+v, want: %+v", test.v, vec, test.want)
				}
			}
		})
	})
}

func TestVector3_Subtract(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		tests := []struct {
			v       *Vector3
			x, y, z float32
			want    *Vector3
		}{
			{NewVector3(0, 0, 0), 1, 2, 3, NewVector3(-1, -2, -3)},
			{NewVector3(1, 1, 1), 11, 11, 11, NewVector3(-10, -10, -10)},
			{NewVector3(16, 16, 16), 6, 1, -1, NewVector3(10, 15, 17)},
			{NewVector3(-3, -5, 3), -3, -5, 3, NewVector3(0, 0, 0)},
			{NewVector3(1e-12, 1e-14, 1e-16), 0, 0, 0, NewVector3(1e-12, 1e-14, 1e-16)},
		}
		t.Run("from XYZ", func(t *testing.T) {
			for _, test := range tests {
				vec := test.v.Subtract(test.x, test.y, test.z)
				if !vec.Equals(test.want) {
					t.Errorf("v: %+v, got: %+v, want: %+v", test.v, vec, test.want)
				}
			}
		})
		t.Run("from Vector", func(t *testing.T) {
			for _, test := range tests {
				vec := test.v.SubtractFromVector3(NewVector3(test.x, test.y, test.z))
				if !vec.Equals(test.want) {
					t.Errorf("v: %+v, got: %+v, want: %+v", test.v, vec, test.want)
				}
			}
		})
	})
}

func TestVector3_Multiply(t *testing.T) {
	tests := []struct {
		v    *Vector3
		num  float32
		want *Vector3
	}{
		{NewVector3(0, 0, 0), 1, NewVector3(0, 0, 0)},
		{NewVector3(math.MaxFloat32, math.MaxFloat32, math.MaxFloat32), 1, NewVector3(math.MaxFloat32, math.MaxFloat32, math.MaxFloat32)},
		{NewVector3(-1, -2, -3), -10, NewVector3(10, 20, 30)},
		{NewVector3(10, -10, 0), 10, NewVector3(100, -100, 0)},
		{NewVector3(1e-12, 1e-14, 1e-16), 2, NewVector3(1e-12*2, 1e-14*2, 1e-16*2)},
	}
	for _, test := range tests {
		got := test.v.Multiply(test.num)
		if !got.Equals(test.want) {
			t.Errorf("vector: %+v, got: %+v, want: %+v", test.v, got, test.want)
		}
	}
}

func TestVector3_Divide(t *testing.T) {
	tests := []struct {
		v    *Vector3
		num  float32
		want *Vector3
	}{
		{NewVector3(0, 0, 0), 1, NewVector3(0, 0, 0)},
		{NewVector3(math.MaxFloat32, math.MaxFloat32, math.MaxFloat32), 2, NewVector3(math.MaxFloat32/2, math.MaxFloat32/2, math.MaxFloat32/2)},
		{NewVector3(-1, -2, -3), -1, NewVector3(1, 2, 3)},
		{NewVector3(10, -10, 0), 10, NewVector3(1, -1, 0)},
		{NewVector3(1e-12, 1e-14, 1e-16), 1, NewVector3(1e-12, 1e-14, 1e-16)},
	}
	for _, test := range tests {
		got := test.v.Divide(test.num)
		if !got.Equals(test.want) {
			t.Errorf("vector: %+v, got: %+v, want: %+v", test.v, got, test.want)
		}
	}
}

func TestVector3_Ceil(t *testing.T) {
	tests := []struct {
		v, want *Vector3
	}{
		{NewVector3(0.1, 1.1, 2.3), NewVector3(1, 2, 3)},
		{NewVector3(math.MaxFloat32, math.MaxFloat32, math.MaxFloat32), NewVector3(math.MaxFloat32, math.MaxFloat32, math.MaxFloat32)},
		{NewVector3(-1.7, -2.3, -3.5), NewVector3(-1, -2, -3)},
		{NewVector3(10.10, -10.0001, 0), NewVector3(11, -10, 0)},
		{NewVector3(-1.14, 5.76, 0.0555), NewVector3(-1, 6, 1)},
	}
	for _, test := range tests {
		got := test.v.Ceil()
		if !got.Equals(test.want) {
			t.Errorf("vector: %+v, got: %+v, want: %+v", test.v, got, test.want)
		}
	}
}

func TestVector3_Floor(t *testing.T) {
	tests := []struct {
		v, want *Vector3
	}{
		{NewVector3(0.1, 1.1, 2.3), NewVector3(0, 1, 2)},
		{NewVector3(math.MaxFloat32, math.MaxFloat32, math.MaxFloat32), NewVector3(math.MaxFloat32, math.MaxFloat32, math.MaxFloat32)},
		{NewVector3(-1.7, -2.3, -3.5), NewVector3(-2, -3, -4)},
		{NewVector3(10.10, -10.0001, 0), NewVector3(10, -11, 0)},
		{NewVector3(-1.14, 5.76, 0.0555), NewVector3(-2, 5, 0)},
	}
	for _, test := range tests {
		got := test.v.Floor()
		if !got.Equals(test.want) {
			t.Errorf("vector: %+v, got: %+v, want: %+v", test.v, got, test.want)
		}
	}
}

func TestVector3_Round(t *testing.T) {
	tests := []struct {
		v, want *Vector3
	}{
		{NewVector3(0.1, 1.1, 2.3), NewVector3(0, 1, 2)},
		{NewVector3(math.MaxFloat32, math.MaxFloat32, math.MaxFloat32), NewVector3(math.MaxFloat32, math.MaxFloat32, math.MaxFloat32)},
		{NewVector3(-1.7, -2.3, -3.5), NewVector3(-2, -2, -4)},
		{NewVector3(10.10, -10.0001, 0), NewVector3(10, -10, 0)},
		{NewVector3(-1.14, 5.76, 0.0555), NewVector3(-1, 6, 0)},
	}
	t.Run("Normal", func(t *testing.T) {
		for _, test := range tests {
			got := test.v.Round()
			if !got.Equals(test.want) {
				t.Errorf("vector: %+v, got: %+v, want: %+v", test.v, got, test.want)
			}
		}
	})
	t.Run("Even", func(t *testing.T) {
		for _, test := range tests {
			got := test.v.RoundToEven()
			if !got.Equals(test.want) {
				t.Errorf("vector: %+v, got: %+v, want: %+v", test.v, got, test.want)
			}
		}
	})
}

func TestVector3_Abs(t *testing.T) {
	tests := []struct {
		v, want *Vector3
	}{
		{NewVector3(0, 0, 0), NewVector3(0, 0, 0)},
		{NewVector3(math.MaxFloat32, -math.MaxFloat32, -math.MaxFloat32), NewVector3(math.MaxFloat32, math.MaxFloat32, math.MaxFloat32)},
		{NewVector3(-1, -2.5, -3), NewVector3(1, 2.5, 3)},
		{NewVector3(10, -10, 0.5), NewVector3(10, 10, 0.5)},
		{NewVector3(-1e-12, 1e-14, -1e-16), NewVector3(1e-12, 1e-14, 1e-16)},
	}
	for _, test := range tests {
		got := test.v.Abs()
		if !got.Equals(test.want) {
			t.Errorf("vector: %+v, got: %+v, want: %+v", test.v, got, test.want)
		}
	}
}

func TestVector3_AsVector3(t *testing.T) {
	tests := []*Vector3{
		NewVector3(0, 0, 0),
		NewVector3(math.MaxFloat32, math.MaxFloat32, math.MaxFloat32),
		NewVector3(-1, -2, -3),
		NewVector3(10, -10, 0),
		NewVector3(1e-12, 1e-14, 1e-16),
	}
	for _, test := range tests {
		vec := test.AsVector3()
		if !test.Equals(vec) {
			t.Errorf("got: %+v, want: %+v", vec, test)
		}
	}
}

func TestVector3_Distance(t *testing.T) {
	type distanceTestCase struct {
		v, pos *Vector3
		want   float32
	}
	t.Run("Squared", func(t *testing.T) {
		tests := []distanceTestCase{
			{NewVector3(1, 1, 1), NewVector3(2, 2, 2), 3},
			{NewVector3(3, 5, 7), NewVector3(1, 3, 5), 12},
			{NewVector3(-1, -2.5, -3), NewVector3(1, 2.5, 3), 65},
			{NewVector3(10, -10, 6), NewVector3(10, 10, 6), 400},
			{NewVector3(-1e-12, 1e-14, -1e-16), NewVector3(1e-12, 1e-14, 1e-16), 4e-24},
		}
		for _, test := range tests {
			got := test.v.DistanceSquared(test.pos)
			if test.want != got {
				t.Errorf("got: %#v, want: %#v", got, test.want)
			}
		}
	})
	t.Run("Normal", func(t *testing.T) {
		tests := []distanceTestCase{
			{NewVector3(1, 1, 1), NewVector3(2, 2, 2), 1.7320508},
			{NewVector3(3, 5, 7), NewVector3(1, 3, 5), 3.4641016},
			{NewVector3(-1, -2.5, -3), NewVector3(1, 2.5, 3), 8.062258},
			{NewVector3(10, -10, 6), NewVector3(10, 10, 6), 20},
			{NewVector3(-1e-12, 1e-14, -1e-16), NewVector3(1e-12, 1e-14, 1e-16), 2e-12},
		}
		for _, test := range tests {
			got := test.v.Distance(test.pos)
			if test.want != got {
				t.Errorf("got: %#v, want: %#v", got, test.want)
			}
		}
	})
}

func TestVector3_MaxPlainDistance(t *testing.T) {
	tests := []struct {
		v, pos *Vector3
		want   float32
	}{
		{NewVector3(0, 0, 0), NewVector3(99, 199, 299), 299},
		{NewVector3(1, 1, 1), NewVector3(2, 2, 2), 1},
		{NewVector3(3, 5, 7), NewVector3(1, 3, 5), 2},
		{NewVector3(-1, -2.5, -3), NewVector3(1, 2.5, 3), 6},
		{NewVector3(10, -10, 6), NewVector3(10, 10, 6), 0},
		{NewVector3(-1e-12, 1e-14, -1e-16), NewVector3(1e-12, 1e-14, 1e-16), 2e-12},
	}
	for _, test := range tests {
		got := test.v.MaxPlainDistance(test.pos)
		if test.want != got {
			t.Errorf("got: %#v, want: %#v", got, test.want)
		}
	}
}

func TestVector3_Length(t *testing.T) {
	type lengthTestCase struct {
		v    *Vector3
		want float32
	}
	t.Run("Squared", func(t *testing.T) {
		tests := []lengthTestCase{
			{NewVector3(3, 5, 7), 83},
			{NewVector3(-1, -2, -3), 14},
			{NewVector3(10, -10, 0), 200},
			{NewVector3(1e-12, 1e-14, 1e-16), 1.0001e-24},
		}
		for _, test := range tests {
			got := test.v.LengthSquared()
			if test.want != got {
				t.Errorf("vector: %+v, got: %#v, want: %#v", test.v, got, test.want)
			}
		}
	})
	t.Run("Normal", func(t *testing.T) {
		tests := []lengthTestCase{
			{NewVector3(3, 5, 7), 9.110434},
			{NewVector3(-1, -2, -3), 3.7416575},
			{NewVector3(10, -10, 0), 14.142136},
			{NewVector3(1e-12, 1e-14, 1e-16), 1.00005e-12},
		}
		for _, test := range tests {
			got := test.v.Length()
			if test.want != got {
				t.Errorf("vector: %#v, got: %#v, want: %#v", test.v, got, test.want)
			}
		}
	})
}

func TestVector3_Normalize(t *testing.T) {
	tests := []struct {
		v, want *Vector3
	}{
		{NewVector3(3, 5, 7), NewVector3(0.32929277, 0.5488213, 0.7683498)},
		{NewVector3(1, 1, 1), NewVector3(0.57735026, 0.57735026, 0.57735026)},
		{NewVector3(16, 16, 16), NewVector3(0.57735026, 0.57735026, 0.57735026)},
		{NewVector3(-3, -5, 3), NewVector3(-0.45749572, -0.7624929, 0.45749572)},
		{NewVector3(1e-12, 1e-14, 1e-16), NewVector3(0.99995, 0.0099995, 9.9995006e-05)},
	}
	for _, test := range tests {
		got := test.v.Normalize()
		if !got.Equals(test.want) {
			t.Errorf("vector: %+v, got: %+v, want: %+v", test.v, got, test.want)
		}
	}
}

func TestVector3_Dot(t *testing.T) {
	tests := []struct {
		v1, v2 *Vector3
		want   float32
	}{
		{NewVector3(1, 1, 1), NewVector3(2, 2, 2), 6},
		{NewVector3(-1, -2, -3), NewVector3(10, -10, 5), -5},
		{NewVector3(10, -10, 10), NewVector3(-10, 10, -10), -300},
		{NewVector3(1e-12, 1e-14, 1e-16), NewVector3(1, 1, 1), 1.0101e-12},
	}
	for _, test := range tests {
		dot := test.v1.Dot(test.v2)
		if test.want != dot {
			t.Errorf("vec1: %+v, vec2: %+v, got: %+v, want: %+v", test.v1, test.v2, dot, test.want)
		}
	}

	// Inf Test
	signedInfVec := NewVector3(math.MaxFloat32, math.MaxFloat32, math.MaxFloat32)
	unsignedInfVec := NewVector3(-math.MaxFloat32, -math.MaxFloat32, -math.MaxFloat32)

	dot := signedInfVec.Dot(signedInfVec)
	if !math.IsInf(float64(dot), 1) {
		t.Errorf("vec: %+v, got: %+v, want: +Inf", signedInfVec, dot)
	}

	dot = unsignedInfVec.Dot(signedInfVec)
	if !math.IsInf(float64(dot), -1) {
		t.Errorf("vec: %+v, got: %+v, want: -Inf", unsignedInfVec, dot)
	}

	dot = unsignedInfVec.Dot(unsignedInfVec)
	if !math.IsInf(float64(dot), 1) {
		t.Errorf("vec: %+v, got: %+v, want: +Inf", signedInfVec, dot)
	}
}

func TestVector3_Cross(t *testing.T) {
	tests := []struct {
		v, cross, want *Vector3
	}{
		{NewVector3(3, 5, 7), NewVector3(1, 2, 3), NewVector3(1, -2, 1)},
		{NewVector3(1, 1, 1), NewVector3(9, 9, 9), NewVector3(0, 0, 0)},
		{NewVector3(16, 16, 16), NewVector3(-6, -11, -13), NewVector3(-32, 112, -80)},
		{NewVector3(-3, -5, 3), NewVector3(3, 5, -3), NewVector3(0, 0, 0)},
		{NewVector3(1e-12, 1e-14, 1e-16), NewVector3(0, 0, 0), NewVector3(0, 0, 0)},
	}
	for _, test := range tests {
		got := test.v.Cross(test.cross)
		if !got.Equals(test.want) {
			t.Errorf("vector: %+v, got: %+v, want: %+v", test.v, got, test.want)
		}
	}
}

func TestVector3_SetComponents(t *testing.T) {
	tests := []struct {
		v, want *Vector3
	}{
		{NewVector3(0, 0, 0), NewVector3(5, 0, -5)},
		{NewVector3(math.MaxFloat32, math.MaxFloat32, math.MaxFloat32), NewVector3(-math.MaxFloat32, -math.MaxFloat32, -math.MaxFloat32)},
		{NewVector3(-1, -2, -3), NewVector3(1, 2, 3)},
		{NewVector3(10, -10, 0), NewVector3(100, 50, 100)},
		{NewVector3(1e-12, 1e-14, 1e-16), NewVector3(1e-12, 1e-14, 1e-16)},
	}
	for _, test := range tests {
		test.v.SetComponents(test.want.X, test.want.Y, test.want.Z)
		if !test.v.Equals(test.want) {
			t.Errorf("got: %+v, want: %+v", test.v, test.want)
		}
	}
}

func TestVector3_GetIntermediateWithValue(t *testing.T) {
	t.Run("X", func(t *testing.T) {
		tests := []struct {
			v1, v2 *Vector3
			x      float32
			want   *Vector3
		}{
			{NewVector3(0, 0, 0), NewVector3(3, 5, 7), 4, nil},
			{NewVector3(10, 10, 10), NewVector3(20, 20, 20), 10, NewVector3(10, 10, 10)},
			{NewVector3(-10, -10, -10), NewVector3(-5, -5, -5), -7, NewVector3(-7, -7, -7)},
			{NewVector3(-32, -16, -32), NewVector3(-16, -32, -16), -128, nil},
			{NewVector3(3, 5, 7), NewVector3(4, 6, 8), 3, NewVector3(3, 5, 7)},
		}
		for _, test := range tests {
			got := test.v1.GetIntermediateWithXValue(test.v2, test.x)
			if test.want == nil {
				if nil != got {
					t.Errorf("vector: %+v, got: %+v, want: %+v", test.v1, got, nil)
				}
			} else {
				if !got.Equals(test.want) {
					t.Errorf("vector: %+v, got: %+v, want: %+v", test.v1, got, test.want)
				}
			}
		}
	})
	t.Run("Y", func(t *testing.T) {
		tests := []struct {
			v1, v2 *Vector3
			y      float32
			want   *Vector3
		}{
			{NewVector3(0, 0, 0), NewVector3(3, 5, 7), 6, nil},
			{NewVector3(10, 10, 10), NewVector3(20, 20, 20), 10, NewVector3(10, 10, 10)},
			{NewVector3(-10, -10, -10), NewVector3(-5, -5, -5), -7, NewVector3(-7, -7, -7)},
			{NewVector3(-32, -16, -32), NewVector3(-16, -32, -16), -128, nil},
			{NewVector3(3, 5, 7), NewVector3(4, 6, 8), 5, NewVector3(3, 5, 7)},
		}
		for _, test := range tests {
			got := test.v1.GetIntermediateWithYValue(test.v2, test.y)
			if test.want == nil {
				if nil != got {
					t.Errorf("vector: %+v, got: %+v, want: %+v", test.v1, got, nil)
				}
			} else {
				if !got.Equals(test.want) {
					t.Errorf("vector: %+v, got: %+v, want: %+v", test.v1, got, test.want)
				}
			}
		}
	})
	t.Run("Z", func(t *testing.T) {
		tests := []struct {
			v1, v2 *Vector3
			z      float32
			want   *Vector3
		}{
			{NewVector3(0, 0, 0), NewVector3(3, 5, 7), 8, nil},
			{NewVector3(10, 10, 10), NewVector3(20, 20, 20), 10, NewVector3(10, 10, 10)},
			{NewVector3(-10, -10, -10), NewVector3(-5, -5, -5), -7, NewVector3(-7, -7, -7)},
			{NewVector3(-32, -16, -32), NewVector3(-16, -32, -16), -128, nil},
			{NewVector3(3, 5, 7), NewVector3(4, 6, 8), 7, NewVector3(3, 5, 7)},
		}
		for _, test := range tests {
			got := test.v1.GetIntermediateWithZValue(test.v2, test.z)
			if test.want == nil {
				if nil != got {
					t.Errorf("vector: %+v, got: %+v, want: %+v", test.v1, got, nil)
				}
			} else {
				if !got.Equals(test.want) {
					t.Errorf("vector: %+v, got: %+v, want: %+v", test.v1, got, test.want)
				}
			}
		}
	})
}
