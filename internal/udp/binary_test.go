package udp

import (
	"math"
	"reflect"
	"testing"
)

// Inspired by Go binary tests: https://golang.org/src/encoding/binary/binary_test.go
type TestPacket struct {
	Int8      int8
	Int16     int16
	Int32     int32
	Int64     int64
	Uint8     uint8
	Uint16    uint16
	Uint32    uint32
	Uint64    uint64
	Float32   float32
	Float64   float64
	Byte      byte
	ByteArray [4]byte
	UintArray [4]uint8
}

var packet = TestPacket{
	Int8:   0x01,
	Int16:  0x0203,
	Int32:  0x04050607,
	Int64:  0x08090a0b0c0d0e0f,
	Uint8:  0x10,
	Uint16: 0x1112,
	Uint32: 0x13141516,
	Uint64: 0x1718191a1b1c1d1e,

	Float32: math.Float32frombits(0x1f202122),
	Float64: math.Float64frombits(0x232425262728292a),

	Byte:      0x2b,
	ByteArray: [4]byte{0x2f, 0x2e, 0x2d, 0x2c},
	UintArray: [4]uint8{0x33, 0x32, 0x31, 0x30},
}

// Little Endian Buffer
var little = []byte{
	1,
	3, 2,
	7, 6, 5, 4,
	15, 14, 13, 12, 11, 10, 9, 8,
	16,
	18, 17,
	22, 21, 20, 19,
	30, 29, 28, 27, 26, 25, 24, 23,

	34, 33, 32, 31,
	42, 41, 40, 39, 38, 37, 36, 35,

	43,
	47, 46, 45, 44,
	51, 50, 49, 48,
}

func TestReadPacket(t *testing.T) {
	p := new(TestPacket)
	err := ReadPacket(little, p)
	if err != nil {
		t.Fatalf("unexpected error: %v: ", err)
	}

	want := &packet
	if !reflect.DeepEqual(p, want) {
		t.Fatalf("got: %v, want: %v", p, want)
	}
}
