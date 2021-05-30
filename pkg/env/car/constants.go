package car

// Fuel Mix
const (
	FuelMixLean     uint8 = 0
	FuelMixStandard uint8 = 1
	FuelMixRich     uint8 = 2
	FuelMixMax      uint8 = 3
)

// ERS Deploy Mode
const (
	ERSDeployModeNone     uint8 = 0
	ERSDeployModeMedium   uint8 = 1
	ERSDeployModeOvertake uint8 = 2
	ERSDeployModeHotLap   uint8 = 3
)

// SLI Pro Support
const (
	SLIProSupportInactive uint8 = 0
	SLIProSupportActive   uint8 = 1
)

// MFD Panel Index
// May vary depending on game mode.
const (
	MFDClosed       uint8 = 255
	MFDCarSetup     uint8 = 0
	MFDPits         uint8 = 1
	MFDDamage       uint8 = 2
	MFDEngine       uint8 = 3
	MFDTemperatures uint8 = 4
)

// Pit Limiter Status
const (
	PitLimiterOff uint8 = 0
	PitLimiterOn  uint8 = 0
)

// Traction Control
const (
	TractionControlOff  uint8 = 0
	TractionControlHigh uint8 = 2
)

// Anti Lock Brakes
const (
	AntiLockBrakesOff uint8 = 0
	AntiLockBrakesOn  uint8 = 1
)
