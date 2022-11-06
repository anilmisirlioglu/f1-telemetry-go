package infringement

// Infringement types
//
//go:generate stringer -type=Infringement -linecomment
type Infringement uint8

const (
	BlockingBySlowDriving                     Infringement = iota // Blocking by slow driving
	BlockingByWrongWayDriving                                     // Blocking by wrong way driving
	ReversingOffTheStartLine                                      // Reversing off the start line
	BigCollision                                                  // Big Collision
	SmallCollision                                                // Small Collision
	CollisionFailedToHandBackPositionSingle                       // Collision failed to hand back position single
	CollisionFailedToHandBackPositionMultiple                     // Collision failed to hand back position multiple
	CornerCuttingGainedTime                                       // Corner cutting gained time
	CornerCuttingOvertakeSingle                                   // Corner cutting overtake single
	CornerCuttingOvertakeMultiple                                 // Corner cutting overtake multiple
	CrossedPitExitLane                                            // Crossed pit exit lane
	IgnoringBlueFlags                                             // Ignoring blue flags
	IgnoringYellowFlags                                           // Ignoring yellow flags
	IgnoringDriveThrough                                          // Ignoring drive through
	TooManyDriveThroughs                                          // Too many drive throughs
	DriveThroughReminderServeWithinNLaps                          // Drive through reminder serve within n laps
	DriveThroughReminderServeThisLap                              // Drive through reminder serve this lap
	PitLaneSpeeding                                               // Pit lane speeding
	ParkedForTooLong                                              // Parked for too long
	IgnoringTyreRegulations                                       // Ignoring tyre regulations
	TooManyPenalties                                              // Too many penalties
	MultipleWarnings                                              // Multiple warnings
	ApproachingDisqualification                                   // Approaching disqualification
	TyreRegulationsSelectSingle                                   // Tyre regulations select single
	TyreRegulationsSelectMultiple                                 // Tyre regulations select multiple
	LapInvalidatedCornerCutting                                   // Lap invalidated corner cutting
	LapInvalidatedRunningWide                                     // Lap invalidated running wide
	CornerCuttingRanWideGainedTimeMinor                           // Corner cutting ran wide gained time minor
	CornerCuttingRanWideGainedTimeSignificant                     // Corner cutting ran wide gained time significant
	CornerCuttingRanWideGainedTimeExtreme                         // Corner cutting ran wide gained time extreme
	LapInvalidatedWallRiding                                      // Lap invalidated wall riding
	LapInvalidatedFlashbackUsed                                   // Lap invalidated flashback used
	LapInvalidatedResetToTrack                                    // Lap invalidated reset to track
	BlockingThePitLane                                            // Blocking the pitlane
	JumpStart                                                     // Jump start
	SafetyCarToCarCollision                                       // Safety car to car collision
	SafetyCarIllegalOvertake                                      // Safety car illegal overtake
	SafetyCarExceedingAllowedPace                                 // Safety car exceeding allowed pace
	VirtualSafetyCarExceedingAllowedPace                          // Virtual safety car exceeding allowed pace
	FormationLapBelowAllowedSpeed                                 // Formation lap below allowed speed
	FormationLapParking                                           // Formation lap parking
	RetiredMechanicalFailure                                      // Retired mechanical failure
	RetiredTerminallyDamaged                                      // Retired terminally damaged
	SafetyCarFallingTooFarBack                                    // Safety car falling too far back
	BlackFlagTimer                                                // Black flag timer
	UnServedStopGoPenalty                                         // Unserved stop go penalty
	UnServedDriveThroughPenalty                                   // Unserved drive through penalty
	EngineComponentChange                                         // Engine component change
	GearboxChange                                                 // Gearbox change
	ParcFermeChange                                               // Parc Fermé change
	LeagueGridPenalty                                             // League grid penalty
	RetryPenalty                                                  // Retry penalty
	IllegalTimeGain                                               // Illegal time gain
	MandatoryPitStop                                              // Mandatory pitstop
	AttributeAssigned                                             // Attribute assigned
)
