package penalty

// Penalty Types
//
//go:generate stringer -type=Penalty -linecomment
type Penalty uint8

const (
	DriveThrough                           Penalty = iota // Drive through
	StopGo                                                // Stop Go
	Grid                                                  // Grid penalty
	Reminder                                              // Penalty reminder
	Time                                                  // Time penalty
	Warning                                               // Warning
	Disqualified                                          // Disqualified
	RemoveFromFormationLap                                // Removed from formation lap
	ParkedTooLongTimer                                    // Parked too long timer
	TypeRegulations                                       // Tyre regulations
	ThisLapInvalidated                                    // This lap invalidated
	ThisAndNextNextLapInvalidated                         // This and next lap invalidated
	ThisLapInvalidatedWithoutReason                       // This lap invalidated without reason
	ThisAndNextLapInvalidatedWithoutReason                // This and next lap invalidated without reason
	ThisAndPrevLapInvalidated                             // This and previous lap invalidated
	ThisAndPrevLapInvalidatedWithoutReason                // This and previous lap invalidated without reason
	Retired                                               // Retired
	BlackFlagTimer                                        // Black flag timer
)
