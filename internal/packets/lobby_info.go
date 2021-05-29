package packets

type LobbyInfoData struct {
	AIControlled uint8    // Whether the vehicle is AI (1) or Human (0) controlled
	TeamID       uint8    // Team ID - see docs/IDS.md#team-ids (255 if no team currently selected)
	Nationality  uint8    // Nationality of the driver
	Name         [48]rune // Name of participant in UTF-8 format – null terminated // Will be truncated with ... (U+2026) if too long
	ReadyStatus  uint8    // 0 = not ready, 1 = ready, 2 = spectating
}

type PacketLobbyInfoData struct {
	Header PacketHeader // Header
	// Packet specific data
	NumPlayers   uint8 // Number of players in the lobby data
	LobbyPlayers [22]LobbyInfoData
}
