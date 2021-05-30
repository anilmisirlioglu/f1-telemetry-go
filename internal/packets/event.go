package packets

// This packet gives details of events that happen during the course of a session.

// Frequency: When the event occurs
// Size: 35 bytes
// Version: 1

type FastestLap struct {
	VehicleIdx uint8   // Vehicle index of car achieving fastest lap
	LapTime    float32 // Lap time is in seconds
}

type Retirement struct {
	VehicleIdx uint8 // Vehicle index of car retiring
}

type TeamMateInPits struct {
	VehicleIdx uint8 // Vehicle index of team mate
}

type RaceWinner struct {
	VehicleIdx uint8 // Vehicle index of the race winner
}

type Penalty struct {
	PenaltyType      uint8 // Penalty type – see docs/TYPES.md#penalty-types
	InfringementType uint8 // Infringement type – see docs/TYPES.md#infringement-types
	VehicleIDx       uint8 // Vehicle index of the car the penalty is applied to
	OtherVehicleIdx  uint8 // Vehicle index of the other car involved
	Time             uint8 // Time gained, or time spent doing action in seconds
	LapNum           uint8 // Lap the penalty occurred on
	PlacesGained     uint8 // Number of places gained by this
}

type SpeedTrap struct {
	VehicleIdx uint8   // Vehicle index of the vehicle triggering speed trap
	Speed      float32 // Top speed achieved in kilometres per hour
}

type EventDataDetails struct {
	FastestLap     FastestLap
	Retirement     Retirement
	TeamMateInPits TeamMateInPits
	RaceWinner     RaceWinner
	Penalty        Penalty
	SpeedTrap      SpeedTrap
}

type PacketEventData struct {
	Header          PacketHeader
	EventStringCode [4]uint8         // Event string code, see below
	EventDetails    EventDataDetails // Event details - should be interpreted differently
	// for each type
}

func (p *PacketEventData) EventCodeString() string {
	b := make([]byte, len(p.EventStringCode))
	for i, v := range p.EventStringCode {
		b[i] = v
	}

	return string(b)
}
