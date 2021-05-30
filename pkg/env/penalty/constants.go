package penalty

// Penalty Types
const (
	DriveThrough                           uint8 = 0
	StopGo                                 uint8 = 1
	Grid                                   uint8 = 2
	Reminder                               uint8 = 3
	Time                                   uint8 = 4
	Warning                                uint8 = 5
	Disqualified                           uint8 = 6
	RemoveFromFormationLap                 uint8 = 7
	ParkedTooLongTimer                     uint8 = 8
	TypeRegulations                        uint8 = 9
	ThisLapInvalidated                     uint8 = 10
	ThisAndNextNextLapInvalidated          uint8 = 11
	ThisLapInvalidatedWithoutReason        uint8 = 12
	ThisAndNextLapInvalidatedWithoutReason uint8 = 13
	ThisAndPrevLapInvalidated              uint8 = 14
	ThisAndPrevLapInvalidatedWithoutReason uint8 = 15
	Retired                                uint8 = 16
	BlackFlagTimer                         uint8 = 17
)
