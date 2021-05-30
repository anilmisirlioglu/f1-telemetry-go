package drs

// DRS Allowed
const (
	NotAllowed uint8 = 0
	Allowed    uint8 = 1
	Unknown    int8  = -1
)

// DRS Fault
const (
	Ok    uint8 = 0
	Fault uint8 = 1
)

// DRS Status
const (
	Off uint8 = 0
	On  uint8 = 1
)
