package track

// Track IDs
type Track uint8

const (
	Melbourne        Track = 0
	PaulRicard       Track = 1
	Shanghai         Track = 2
	Sakhir           Track = 3
	Catalunya        Track = 4
	Monaco           Track = 5
	Montreal         Track = 6
	Silverstone      Track = 7
	Hockenheim       Track = 8
	Hungaroring      Track = 9
	Spa              Track = 10
	Monza            Track = 11
	Singapore        Track = 12
	Suzuka           Track = 13
	AbuDhabi         Track = 14
	Texas            Track = 15
	Brazil           Track = 16
	Austria          Track = 17
	Sochi            Track = 18
	Mexico           Track = 19
	Baku             Track = 20
	SakhirShort      Track = 21
	SilverstoneShort Track = 22
	TexasShort       Track = 23
	SuzukaShort      Track = 24
	Hanoi            Track = 25
	Zandvoort        Track = 26
	Imola            Track = 27
	Partimao         Track = 28
	Jeddah           Track = 29
)

var tracks = map[Track]string{
	0:  "Melbourne",
	1:  "PaulRicard",
	2:  "Shanghai",
	3:  "Sakhir",
	4:  "Catalunya",
	5:  "Monaco",
	6:  "Montreal",
	7:  "Silverstone",
	8:  "Hockenheim",
	9:  "Hungaroring",
	10: "Spa",
	11: "Monza",
	12: "Singapore",
	13: "Suzuka",
	14: "AbuDhabi",
	15: "Texas",
	16: "Brazil",
	17: "Austria",
	18: "Sochi",
	19: "Mexico",
	20: "Baku",
	21: "SakhirShort",
	22: "SilverstoneShort",
	23: "TexasShort",
	24: "SuzukaShort",
	25: "Hanoi",
	26: "Zandvoort",
	27: "Imola",
	28: "Partimao",
	29: "Jeddah",
}

func (t Track) String() string {
	return tracks[t]
}
