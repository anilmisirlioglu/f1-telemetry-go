package env

// Packet IDs
const (
	PacketMotion              uint8 = 0  // Contains all motion data for player’s car – only sent while player is in control
	PacketSession             uint8 = 1  // Data about the session – track, time left
	PacketLap                 uint8 = 2  // Data about all the lap times of cars in the session
	PacketEvent               uint8 = 3  // Various notable events that happen during a session
	PacketParticipants        uint8 = 4  // List of participants in the session, mostly relevant for multiplayer
	PacketCarSetup            uint8 = 5  // Packet detailing car setups for cars in the race
	PacketCarTelemetry        uint8 = 6  // Telemetry data for all cars
	PacketCarStatus           uint8 = 7  // Status data for all cars such as damage
	PacketFinalClassification uint8 = 8  // Final classification confirmation at the end of a race
	PacketLobbyInfo           uint8 = 9  // Information about players in a multiplayer lobby
	PacketCarDamage           uint8 = 10 // Damage status for all cars
	PacketSessionHistory      uint8 = 11 // Lap and tyre data for session
)
