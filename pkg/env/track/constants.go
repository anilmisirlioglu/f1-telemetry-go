package track

// Track IDs

//go:generate stringer -type=Track -linecomment
type Track uint8

const (
	Melbourne  Track = iota
	PaulRicard       // Paul Ricard
	Shanghai
	Sakhir
	Catalunya
	Monaco
	Montreal
	Silverstone
	Hockenheim
	Hungaroring
	Spa
	Monza
	Singapore
	Suzuka
	AbuDhabi // Abu Dhabi
	Texas
	Brazil
	Austria
	Sochi
	Mexico
	Baku
	SakhirShort      // Sakhir Short
	SilverstoneShort // Silverstone Short
	TexasShort       // Texas Short
	SuzukaShort      // Suzuka Short
	Hanoi
	Zandvoort
	Imola
	Partimao
	Jeddah
	Miami
)
