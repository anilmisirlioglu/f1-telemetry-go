package team

// Team IDs
type TeamsType uint8

const (
	Mercedes                TeamsType = 0
	Ferrari                 TeamsType = 1
	RedBulRacing            TeamsType = 2
	Williams                TeamsType = 3
	AstonMartin             TeamsType = 4
	Alpine                  TeamsType = 5
	AlphaTauri              TeamsType = 6
	Haas                    TeamsType = 7
	McLaren                 TeamsType = 8
	AlfaRomeo               TeamsType = 9
	McLaren1988             TeamsType = 10
	McLaren1991             TeamsType = 11
	Williams1992            TeamsType = 12
	Ferrari1995             TeamsType = 13
	Williams1996            TeamsType = 14
	McLaren1998             TeamsType = 15
	Ferrari2002             TeamsType = 16
	Ferrari2004             TeamsType = 17
	Renault2006             TeamsType = 18
	Ferrari2007             TeamsType = 19
	McLaren2008             TeamsType = 20
	RedBull2010             TeamsType = 21
	Ferrari1976             TeamsType = 22
	ARTGrandPrix            TeamsType = 23
	CamposVexatecRacing     TeamsType = 24
	Carlin                  TeamsType = 25
	CharouzRacingSystem     TeamsType = 26
	DAMS                    TeamsType = 27
	RussianTime             TeamsType = 28
	MPMotorsport            TeamsType = 29
	Pertamina               TeamsType = 30
	McLaren1990             TeamsType = 31
	Trident                 TeamsType = 32
	BWTArden                TeamsType = 33
	McLaren1976             TeamsType = 34
	Lotus1972               TeamsType = 35
	Ferrari1979             TeamsType = 36
	McLaren1982             TeamsType = 37
	Williams2003            TeamsType = 38
	Brawn2009               TeamsType = 39
	Lotus1978               TeamsType = 40
	F1GenericCar            TeamsType = 41
	ArtGP2019               TeamsType = 42
	Campos2019              TeamsType = 43
	Carlin2019              TeamsType = 44
	SauberJuniorCharouz2019 TeamsType = 45
	Dams2019                TeamsType = 46
	UniVirtuosi2019         TeamsType = 47
	MPMotorsport2019        TeamsType = 48
	Prema2019               TeamsType = 49
	Trident2019             TeamsType = 50
	Arden2019               TeamsType = 51
	Benetton1994            TeamsType = 53
	Benetton1995            TeamsType = 54
	Ferrari2000             TeamsType = 55
	Jordan1991              TeamsType = 56
	Ferrari1990             TeamsType = 63
	McLaren2010             TeamsType = 64
	Ferrari2010             TeamsType = 64
	ArtGP2020               TeamsType = 70
	Campos2020              TeamsType = 71
	Carlin2020              TeamsType = 72
	Charouz2020             TeamsType = 73
	Dam2020                 TeamsType = 74
	UniVirtuosi2020         TeamsType = 75
	MPMotorsport2020        TeamsType = 76
	Prema2020               TeamsType = 77
	Trident2020             TeamsType = 78
	BWT2020                 TeamsType = 79
	Hitech2020              TeamsType = 80
	Mercedes2020            TeamsType = 85
	Ferrari2020             TeamsType = 86
	RedBull2020             TeamsType = 87
	Williams2020            TeamsType = 88
	RacingPoint2020         TeamsType = 89
	Renault2020             TeamsType = 90
	AlphaTauri2020          TeamsType = 91
	Haas2020                TeamsType = 92
	McLaren2020             TeamsType = 93
	AlfaRomeo2020           TeamsType = 94
	Prema2021               TeamsType = 106
	UniVirtuosi2021         TeamsType = 107
	Carlin2021              TeamsType = 108
	Hitech2021              TeamsType = 109
	ArtGP2021               TeamsType = 110
	MPMotorsport2021        TeamsType = 111
	Charouz2021             TeamsType = 112
	Dams2021                TeamsType = 113
	Campos2021              TeamsType = 114
	BWT2021                 TeamsType = 115
	Trident2021             TeamsType = 116
	MyTeam                  TeamsType = 255
)

var teamNames = map[TeamsType]string{
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

func (s TeamsType) String() string {
	return teamNames[s]
}
