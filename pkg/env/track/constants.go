package track

// Track IDs
type TrackType uint8

const (
	Melbourne        TrackType = 0
	PaulRicard       TrackType = 1
	Shanghai         TrackType = 2
	Sakhir           TrackType = 3
	Catalunya        TrackType = 4
	Monaco           TrackType = 5
	Montreal         TrackType = 6
	Silverstone      TrackType = 7
	Hockenheim       TrackType = 8
	Hungaroring      TrackType = 9
	Spa              TrackType = 10
	Monza            TrackType = 11
	Singapore        TrackType = 12
	Suzuka           TrackType = 13
	AbuDhabi         TrackType = 14
	Texas            TrackType = 15
	Brazil           TrackType = 16
	Austria          TrackType = 17
	Sochi            TrackType = 18
	Mexico           TrackType = 19
	Baku             TrackType = 20
	SakhirShort      TrackType = 21
	SilverstoneShort TrackType = 22
	TexasShort       TrackType = 23
	SuzukaShort      TrackType = 24
	Hanoi            TrackType = 25
	Zandvoort        TrackType = 26
	Imola            TrackType = 27
	Partimao         TrackType = 28
	Jeddah           TrackType = 29
)

var tracks = map[TrackType]string{
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

func (t TrackType) String() string {
	return tracks[t]
}
