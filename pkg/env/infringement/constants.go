package infringement

// Infringement types
const (
	BlockingBySlowDriving                     uint8 = 0
	BlockingByWrongWayDriving                 uint8 = 1
	ReversingOffTheStartLine                  uint8 = 2
	BigCollision                              uint8 = 3
	SmallCollision                            uint8 = 4
	CollisionFailedToHandBackPositionSingle   uint8 = 5
	CollisionFailedToHandBackPositionMultiple uint8 = 6
	CornerCuttingGainedTime                   uint8 = 7
	CornerCuttingOvertakeSingle               uint8 = 8
	CornerCuttingOvertakeMultiple             uint8 = 9
	CrossedPitExitLane                        uint8 = 10
	IgnoringBlueFlags                         uint8 = 11
	IgnoringYellowFlags                       uint8 = 12
	IgnoringDriveThrough                      uint8 = 13
	TooManyDriveThroughs                      uint8 = 14
	DriveThroughReminderServeWithinNLaps      uint8 = 15
	DriveThroughReminderServeThisLap          uint8 = 16
	PitLaneSpeeding                           uint8 = 17
	ParkedForTooLong                          uint8 = 18
	IgnoringTyreRegulations                   uint8 = 19
	TooManyPenalties                          uint8 = 20
	MultipleWarnings                          uint8 = 21
	ApproachingDisqualification               uint8 = 22
	TyreRegulationsSelectSingle               uint8 = 23
	TyreRegulationsSelectMultiple             uint8 = 24
	LapInvalidatedCornerCutting               uint8 = 25
	LapInvalidatedRunningWide                 uint8 = 26
	CornerCuttingRanWideGainedTimeMinor       uint8 = 27
	CornerCuttingRanWideGainedTimeSignificant uint8 = 28
	CornerCuttingRanWideGainedTimeExtreme     uint8 = 29
	LapInvalidatedWallRiding                  uint8 = 30
	LapInvalidatedFlashbackUsed               uint8 = 31
	LapInvalidatedResetToTrack                uint8 = 32
	BlockingThePitLane                        uint8 = 33
	JumpStart                                 uint8 = 34
	SafetyCarToCarCollision                   uint8 = 35
	SafetyCarIllegalOvertake                  uint8 = 36
	SafetyCarExceedingAllowedPace             uint8 = 37
	VirtualSafetyCarExceedingAllowedPace      uint8 = 38
	FormationLapBelowAllowedSpeed             uint8 = 39
	RetiredMechanicalFailure                  uint8 = 40
	RetiredTerminallyDamaged                  uint8 = 41
	SafetyCarFallingTooFarBack                uint8 = 42
	BlackFlagTimer                            uint8 = 43
	UnServedStopGoPenalty                     uint8 = 44
	UnServedDriveThroughPenalty               uint8 = 45
	EngineComponentChange                     uint8 = 46
	GearboxChange                             uint8 = 47
	LeagueGridPenalty                         uint8 = 48
	RetryPenalty                              uint8 = 49
	IllegalTimeGain                           uint8 = 50
	MandatoryPitStop                          uint8 = 51
)
