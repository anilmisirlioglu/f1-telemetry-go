package math

import "math"

type Vector3 struct {
	X, Y, Z float32
}

func NewVector3(x, y, z float32) *Vector3 {
	return &Vector3{X: x, Y: y, Z: z}
}

func (v *Vector3) Add(x, y, z float32) *Vector3 {
	return NewVector3(v.X+x, v.Y+y, v.Z+z)
}

func (v *Vector3) AddFromVector3(vector3 *Vector3) *Vector3 {
	return NewVector3(v.X+vector3.X, v.Y+vector3.Y, v.Z+vector3.Z)
}

func (v *Vector3) Subtract(x, y, z float32) *Vector3 {
	return v.Add(-x, -y, -z)
}

func (v *Vector3) SubtractFromVector3(vector3 *Vector3) *Vector3 {
	return v.Add(-vector3.X, -vector3.Y, -vector3.Z)
}

func (v *Vector3) Multiply(num float32) *Vector3 {
	return NewVector3(v.X*num, v.Y*num, v.Z*num)
}

func (v *Vector3) Divide(num float32) *Vector3 {
	return NewVector3(v.X/num, v.Y/num, v.Z/num)
}

func (v *Vector3) Ceil() *Vector3 {
	return NewVector3(
		float32(math.Ceil(float64(v.X))),
		float32(math.Ceil(float64(v.Y))),
		float32(math.Ceil(float64(v.Z))),
	)
}

func (v *Vector3) Floor() *Vector3 {
	return NewVector3(
		float32(math.Floor(float64(v.X))),
		float32(math.Floor(float64(v.Y))),
		float32(math.Floor(float64(v.Z))),
	)
}

func (v *Vector3) Round() *Vector3 {
	return NewVector3(
		float32(math.Round(float64(v.X))),
		float32(math.Round(float64(v.Y))),
		float32(math.Round(float64(v.Z))),
	)
}

func (v *Vector3) RoundToEven() *Vector3 {
	return NewVector3(
		float32(math.RoundToEven(float64(v.X))),
		float32(math.RoundToEven(float64(v.Y))),
		float32(math.RoundToEven(float64(v.Z))),
	)
}

func (v *Vector3) Abs() *Vector3 {
	return NewVector3(
		float32(math.Abs(float64(v.X))),
		float32(math.Abs(float64(v.Y))),
		float32(math.Abs(float64(v.Z))),
	)
}

func (v *Vector3) AsVector3() *Vector3 {
	return NewVector3(v.X, v.Y, v.Z)
}

func (v *Vector3) Distance(pos *Vector3) float32 {
	return sqrt32(v.DistanceSquared(pos))
}

func (v *Vector3) DistanceSquared(pos *Vector3) float32 {
	return pow32(v.X, pos.X, 2) + pow32(v.Y, pos.Y, 2) + pow32(v.Z, pos.Z, 2)
}

func (v *Vector3) MaxPlainDistance(pos *Vector3) float32 {
	return float32(math.Max(
		math.Abs(float64(v.X-pos.X)),
		math.Abs(float64(v.Z-pos.Z)),
	))
}

func (v *Vector3) Length() float32 {
	return sqrt32(v.LengthSquared())
}

func (v *Vector3) LengthSquared() float32 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v *Vector3) Normalize() *Vector3 {
	l := v.LengthSquared()
	if l > 0 {
		return v.Divide(sqrt32(l))
	}

	return new(Vector3)
}

func (v *Vector3) Dot(pos *Vector3) float32 {
	return v.X*pos.X + v.Y*pos.Y + v.Z*pos.Z
}

func (v *Vector3) Cross(vector3 *Vector3) *Vector3 {
	return NewVector3(
		v.Y*vector3.Z-v.Z*vector3.Y,
		v.Z*vector3.X-v.X*vector3.Z,
		v.X*vector3.Y-v.Y*vector3.X,
	)
}

func (v *Vector3) Equals(vector3 *Vector3) bool {
	return v.X == vector3.X && v.Y == vector3.Y && v.Z == vector3.Z
}

func (v *Vector3) SetComponents(x, y, z float32) {
	v.X = x
	v.Y = y
	v.Z = z
}

// GetIntermediateWithXValue returns a new vector with x value equal to the second parameter,
// along the line between this vector and the passed in vector, or nil if not possible.
func (v *Vector3) GetIntermediateWithXValue(vector3 *Vector3, x float32) *Vector3 {
	xDiff := vector3.X - v.X
	yDiff := vector3.Y - v.Y
	zDiff := vector3.Z - v.Z

	if xDiff*xDiff < 0.0000001 {
		return nil
	}

	f := (x - v.X) / xDiff
	if f < 0 || f > 1 {
		return nil
	} else {
		return NewVector3(x, v.Y+yDiff*f, v.Z+zDiff*f)
	}
}

// GetIntermediateWithYValue returns a new vector with y value equal to the second parameter,
// along the line between this vector and the passed in vector, or nil if not possible.
func (v *Vector3) GetIntermediateWithYValue(vector3 *Vector3, y float32) *Vector3 {
	xDiff := vector3.X - v.X
	yDiff := vector3.Y - v.Y
	zDiff := vector3.Z - v.Z

	if yDiff*yDiff < 0.0000001 {
		return nil
	}

	f := (y - v.Y) / yDiff
	if f < 0 || f > 1 {
		return nil
	} else {
		return NewVector3(v.X+xDiff*f, y, v.Z+zDiff*f)
	}
}

// GetIntermediateWithZValue returns a new vector with z value equal to the second parameter,
// along the line between this vector and the passed in vector, or nil if not possible.
func (v *Vector3) GetIntermediateWithZValue(vector3 *Vector3, z float32) *Vector3 {
	xDiff := vector3.X - v.X
	yDiff := vector3.Y - v.Y
	zDiff := vector3.Z - v.Z

	if zDiff*zDiff < 0.0000001 {
		return nil
	}

	f := (z - v.Z) / zDiff
	if f < 0 || f > 1 {
		return nil
	} else {
		return NewVector3(v.X+xDiff*f, v.Y+yDiff*f, z)
	}
}
