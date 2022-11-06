package session

// Type Session Types
//
//go:generate stringer -type=Type -linecomment
type Type uint8

const (
	Unknown Type = iota
	P1
	P2
	P3
	ShortP // Short P
	Q1
	Q2
	Q3
	ShortQ // Short Q
	OSQ
	R
	R2
	R3
	TimeTrial // Time Trial
)

// Length Type lengths
type Length uint8

const (
	None Length = iota
	_
	VeryShort
	Short
	Medium
	MediumLong
	Long
	Full
)
