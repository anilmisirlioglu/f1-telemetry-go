package status

// Result Status
const (
	ResultInvalid       uint8 = 0
	ResultInactive      uint8 = 1
	ResultActive        uint8 = 2
	ResultFinished      uint8 = 3
	ResultDisqualified  uint8 = 4
	ResultNotClassified uint8 = 5
	ResultRetired       uint8 = 6
)

// Driver Status
const (
	DriverInGarage  uint8 = 0
	DriverFlyingLap uint8 = 1
	DriverInLap     uint8 = 2
	DriverOutLap    uint8 = 3
	DriverOnTrack   uint8 = 4
)

// Pit Status
const (
	PitNone      uint8 = 0
	PitPitting   uint8 = 1
	PitInPitArea uint8 = 2
)

// Safety Car Status
const (
	SafetyCarNo      uint8 = 0
	SafetyCarFull    uint8 = 1
	SafetyCarVirtual uint8 = 2
)

// Ready Status
const (
	NotReady   uint8 = 0
	Ready      uint8 = 1
	Spectating uint8 = 2
)
