package game

// Formula
const (
	F1Modern  uint8 = 0
	F1Classic uint8 = 1
	F2        uint8 = 2
	F1Generic uint8 = 3
)

// Sectors
const (
	Sector1 uint8 = 0
	Sector2 uint8 = 1
	Sector3 uint8 = 2
)

// Network Game
const (
	Offline uint8 = 0
	Online  uint8 = 1
)

// UDP Settings
const (
	UDPRestricted uint8 = 0
	UDPPublic     uint8 = 1
)

// Current Lap Invalid
const (
	LapValid   uint8 = 0
	LapInvalid uint8 = 1
)

// AI Controlled
const (
	ControlledHuman uint8 = 0
	ControlledAI    uint8 = 1
)
