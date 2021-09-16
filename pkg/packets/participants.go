package packets

// This is a list of participants in the race. If the vehicle is controlled by AI, then the name will be the driver name.
// If this is a multiplayer game, the names will be the Steam Id on PC, or the LAN name if appropriate.

// Frequency: Every 5 seconds
// Size: 1213 bytes
// Version: 1

type ParticipantData struct {
	AIControlled  uint8    // Whether the vehicle is AI (1) or Human (0) controlled
	DriverID      uint8    // Driver id - see appendix
	TeamID        uint8    // Team id - see appendix
	MyTeam        uint8    // My team flag – 1 = My Team, 0 = otherwise
	RaceNumber    uint8    // Race number of the car
	Nationality   uint8    // Nationality of the driver
	Name          [48]byte // Name of participant in UTF-8 format – null terminated, Will be truncated with … (U+2026) if too long
	YourTelemetry uint8    // The player's UDP setting, 0 = restricted, 1 = public
}

type PacketParticipantsData struct {
	Header        PacketHeader
	NumActiveCars uint8               // Number of active cars in the data – should match number of
	Participants  [22]ParticipantData // cars on HUD
}

func (p *ParticipantData) NameToString() string {
	return string(p.Name[:])
}
