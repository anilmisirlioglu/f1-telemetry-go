package team

// Team IDs
type Team uint8

const (
	Mercedes                Team = 0
	Ferrari                 Team = 1
	RedBulRacing            Team = 2
	Williams                Team = 3
	AstonMartin             Team = 4
	Alpine                  Team = 5
	AlphaTauri              Team = 6
	Haas                    Team = 7
	McLaren                 Team = 8
	AlfaRomeo               Team = 9
	McLaren1988             Team = 10
	McLaren1991             Team = 11
	Williams1992            Team = 12
	Ferrari1995             Team = 13
	Williams1996            Team = 14
	McLaren1998             Team = 15
	Ferrari2002             Team = 16
	Ferrari2004             Team = 17
	Renault2006             Team = 18
	Ferrari2007             Team = 19
	McLaren2008             Team = 20
	RedBull2010             Team = 21
	Ferrari1976             Team = 22
	ARTGrandPrix            Team = 23
	CamposVexatecRacing     Team = 24
	Carlin                  Team = 25
	CharouzRacingSystem     Team = 26
	DAMS                    Team = 27
	RussianTime             Team = 28
	MPMotorsport            Team = 29
	Pertamina               Team = 30
	McLaren1990             Team = 31
	Trident                 Team = 32
	BWTArden                Team = 33
	McLaren1976             Team = 34
	Lotus1972               Team = 35
	Ferrari1979             Team = 36
	McLaren1982             Team = 37
	Williams2003            Team = 38
	Brawn2009               Team = 39
	Lotus1978               Team = 40
	F1GenericCar            Team = 41
	ArtGP2019               Team = 42
	Campos2019              Team = 43
	Carlin2019              Team = 44
	SauberJuniorCharouz2019 Team = 45
	Dams2019                Team = 46
	UniVirtuosi2019         Team = 47
	MPMotorsport2019        Team = 48
	Prema2019               Team = 49
	Trident2019             Team = 50
	Arden2019               Team = 51
	Benetton1994            Team = 53
	Benetton1995            Team = 54
	Ferrari2000             Team = 55
	Jordan1991              Team = 56
	Ferrari1990             Team = 63
	McLaren2010             Team = 64
	Ferrari2010             Team = 64
	ArtGP2020               Team = 70
	Campos2020              Team = 71
	Carlin2020              Team = 72
	Charouz2020             Team = 73
	Dam2020                 Team = 74
	UniVirtuosi2020         Team = 75
	MPMotorsport2020        Team = 76
	Prema2020               Team = 77
	Trident2020             Team = 78
	BWT2020                 Team = 79
	Hitech2020              Team = 80
	Mercedes2020            Team = 85
	Ferrari2020             Team = 86
	RedBull2020             Team = 87
	Williams2020            Team = 88
	RacingPoint2020         Team = 89
	Renault2020             Team = 90
	AlphaTauri2020          Team = 91
	Haas2020                Team = 92
	McLaren2020             Team = 93
	AlfaRomeo2020           Team = 94
	Prema2021               Team = 106
	UniVirtuosi2021         Team = 107
	Carlin2021              Team = 108
	Hitech2021              Team = 109
	ArtGP2021               Team = 110
	MPMotorsport2021        Team = 111
	Charouz2021             Team = 112
	Dams2021                Team = 113
	Campos2021              Team = 114
	BWT2021                 Team = 115
	Trident2021             Team = 116
	MyTeam                  Team = 255
)

var teams = map[Team]string{
	0:   "Mercedes",
	1:   "Ferrari",
	2:   "RedBulRacing",
	3:   "Williams",
	4:   "AstonMartin",
	5:   "Alpine",
	6:   "AlphaTauri",
	7:   "Haas",
	8:   "McLaren",
	9:   "AlfaRomeo",
	10:  "McLaren1988",
	11:  "McLaren1991",
	12:  "Williams1992",
	13:  "Ferrari1995",
	14:  "Williams1996",
	15:  "McLaren1998",
	16:  "Ferrari2002",
	17:  "Ferrari2004",
	18:  "Renault2006",
	19:  "Ferrari2007",
	20:  "McLaren2008",
	21:  "RedBull2010",
	22:  "Ferrari1976",
	23:  "ARTGrandPrix",
	24:  "CamposVexatecRacing",
	25:  "Carlin",
	26:  "CharouzRacingSystem",
	27:  "DAMS",
	28:  "RussianTime",
	29:  "MPMotorsport",
	30:  "Pertamina",
	31:  "McLaren1990",
	32:  "Trident",
	33:  "BWTArden",
	34:  "McLaren1976",
	35:  "Lotus1972",
	36:  "Ferrari1979",
	37:  "McLaren1982",
	38:  "Williams2003",
	39:  "Brawn2009",
	40:  "Lotus1978",
	41:  "F1GenericCar",
	42:  "ArtGP2019",
	43:  "Campos2019",
	44:  "Carlin2019",
	45:  "SauberJuniorCharouz2019",
	46:  "Dams2019",
	47:  "UniVirtuosi2019",
	48:  "MPMotorsport2019",
	49:  "Prema2019",
	50:  "Trident2019",
	51:  "Arden2019",
	53:  "Benetton1994",
	54:  "Benetton1995",
	55:  "Ferrari2000",
	56:  "Jordan1991",
	63:  "Ferrari1990",
	64:  "McLaren2010",
	65:  "Ferrari2010", // in enum 64; don't know what's correct
	70:  "ArtGP2020",
	71:  "Campos2020",
	72:  "Carlin2020",
	73:  "Charouz2020",
	74:  "Dam2020",
	75:  "UniVirtuosi2020",
	76:  "MPMotorsport2020",
	77:  "Prema2020",
	78:  "Trident2020",
	79:  "BWT2020",
	80:  "Hitech2020",
	85:  "Mercedes2020",
	86:  "Ferrari2020",
	87:  "RedBull2020",
	88:  "Williams2020",
	89:  "RacingPoint2020",
	90:  "Renault2020",
	91:  "AlphaTauri2020",
	92:  "Haas2020",
	93:  "McLaren2020",
	94:  "AlfaRomeo2020",
	106: "Prema2021",
	107: "UniVirtuosi2021",
	108: "Carlin2021",
	109: "Hitech2021",
	110: "ArtGP2021",
	111: "MPMotorsport2021",
	112: "Charouz2021",
	113: "Dams2021",
	114: "Campos2021",
	115: "BWT2021",
	116: "Trident2021",
	255: "MyTeam",
}

func (s Team) String() string {
	return teams[s]
}
