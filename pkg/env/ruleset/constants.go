package ruleset

// Ruleset IDS
//
//go:generate stringer -type=Ruleset -linecomment
type Ruleset uint8

const (
	PracticeAndQualifying Ruleset = 0  // Practice & Qualifying
	Race                  Ruleset = 1  // Race
	TimeTrial             Ruleset = 2  // Time Trial
	TimeAttack            Ruleset = 4  // Time Attack
	CheckpointChallenge   Ruleset = 6  // Checkpoint Challenge
	Autocross             Ruleset = 8  // Autocross
	Drift                 Ruleset = 9  // Drift
	AverageSpeedZone      Ruleset = 10 // Average Speed Zone
	RivalDuel             Ruleset = 11 // Rival Duel
)
