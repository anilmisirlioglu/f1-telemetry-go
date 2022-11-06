package game

// Mode IDs
//
//go:generate stringer -type=Mode -linecomment
type Mode uint8

const (
	EventMode                Mode = 0   // Event Mode
	GrandPrix                Mode = 3   // Grand Prix
	TimeTrial                Mode = 5   // Time Trial
	Splitscreen              Mode = 6   // Splitscreen
	OnlineCustom             Mode = 7   // Online Custom
	OnlineLeague             Mode = 8   // Online League
	CareerInvitational       Mode = 11  // Career Invitational
	ChampionshipInvitational Mode = 12  // Championship Invitational
	Championship             Mode = 13  // Championship
	OnlineChampionship       Mode = 14  // Online Championship
	OnlineWeeklyEvent        Mode = 15  // Online Weekly Event
	Career22                 Mode = 19  // Career ‘22
	Career22Online           Mode = 20  // Career ’22 Online
	Benchmark                Mode = 127 // Benchmark
)

// Formula
const (
	F1Modern uint8 = iota
	F1Classic
	F2
	F1Generic
	Beta
	Supercars
	ESports
	F22021
)

// Sectors
const (
	Sector1 uint8 = 0
	Sector2 uint8 = 1
	Sector3 uint8 = 2
)

// Network Game
const (
	Offline uint8 = 0
	Online  uint8 = 1
)

// UDP Settings
const (
	UDPRestricted uint8 = 0
	UDPPublic     uint8 = 1
)

// Current Lap Invalid
const (
	LapValid   uint8 = 0
	LapInvalid uint8 = 1
)

// AI Controlled
const (
	ControlledHuman uint8 = 0
	ControlledAI    uint8 = 1
)
