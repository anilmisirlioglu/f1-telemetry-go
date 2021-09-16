package packets

// This packet details the players currently in a multiplayer lobby.
// It details each player’s selected car, any AI involved in the game and also the ready status of each of the participants.

// Frequency: Two every second when in the lobby
// Size: 1169 bytes
// Version: 1

type LobbyInfoData struct {
	AIControlled uint8    // Whether the vehicle is AI (1) or Human (0) controlled
	TeamID       uint8    // Team ID - see docs/IDS.md#team-ids (255 if no team currently selected)
	Nationality  uint8    // Nationality of the driver
	Name         [48]byte // Name of participant in UTF-8 format – null terminated // Will be truncated with ... (U+2026) if too long
	CarNumber    uint8    // Car number of the player
	ReadyStatus  uint8    // 0 = not ready, 1 = ready, 2 = spectating
}

type PacketLobbyInfoData struct {
	Header PacketHeader // Header
	// Packet specific data
	NumPlayers   uint8 // Number of players in the lobby data
	LobbyPlayers [22]LobbyInfoData
}

func (p *LobbyInfoData) NameToString() string {
	return string(p.Name[:])
}
