package driver

// Drivers IDs
type DriversType uint8

const (
	CarlosSainz         DriversType = 0
	DaniilKvyat         DriversType = 1
	DanielRicciardo     DriversType = 2
	FernandoAlonso      DriversType = 3
	FelipeMassa         DriversType = 4
	KimiRaikkonen       DriversType = 6
	LewisHamilton       DriversType = 7
	MaxVerstappen       DriversType = 9
	NicoHulkenburg      DriversType = 10
	KevinMagnussen      DriversType = 11
	RomainGrosjean      DriversType = 12
	SebastianVettel     DriversType = 13
	SergioPerez         DriversType = 14
	ValtteriBottas      DriversType = 15
	EstebanOcon         DriversType = 17
	LanceStroll         DriversType = 19
	ArronBarnes         DriversType = 20
	MartinGiles         DriversType = 21
	AlexMurray          DriversType = 22
	LucasRoth           DriversType = 23
	IgorCorreia         DriversType = 24
	SophieLevasseur     DriversType = 25
	JonasSchiffer       DriversType = 26
	AlainForest         DriversType = 27
	JayLetourneau       DriversType = 28
	EstoSaari           DriversType = 29
	YasarAtiyeh         DriversType = 30
	CallistoCalabresi   DriversType = 31
	NaotaIzum           DriversType = 32
	HowardClarke        DriversType = 33
	WilheimKaufmann     DriversType = 34
	MarieLaursen        DriversType = 35
	FlavioNieves        DriversType = 36
	PeterBelousov       DriversType = 37
	KlimekMichalski     DriversType = 38
	SantiagoMoreno      DriversType = 39
	BenjaminCoppens     DriversType = 40
	NoahVisser          DriversType = 41
	GertWaldmuller      DriversType = 42
	JulianQuesada       DriversType = 43
	DanielJones         DriversType = 44
	ArtemMarkelov       DriversType = 45
	TadasukeMakino      DriversType = 46
	SeanGelael          DriversType = 47
	NyckDeVries         DriversType = 48
	JackAitken          DriversType = 49
	GeorgeRussell       DriversType = 50
	MaximilianGunther   DriversType = 51
	NireiFukuzumi       DriversType = 52
	LucaGhiotto         DriversType = 53
	LandoNorris         DriversType = 54
	SergioSetteCamara   DriversType = 55
	LouisDeletraz       DriversType = 56
	AntonioFuoco        DriversType = 57
	CharlesLeclerc      DriversType = 58
	PierreGasly         DriversType = 59
	AlexanderAlbon      DriversType = 62
	NicholasLatifi      DriversType = 63
	DorianBoccolacci    DriversType = 64
	NikoKari            DriversType = 65
	RobertoMerhi        DriversType = 66
	ArjunMaini          DriversType = 67
	AlessioLorandi      DriversType = 68
	RubenMeijer         DriversType = 69
	RashidNair          DriversType = 70
	JackTremblay        DriversType = 71
	AntonioGiovinazzi   DriversType = 74
	RobertKubica        DriversType = 75
	NobuharuMatsushita  DriversType = 78
	NikitaMazepin       DriversType = 79
	GuanyaZhou          DriversType = 80
	MickSchumacher      DriversType = 81
	CallumIlott         DriversType = 82
	JuanManuelCorrea    DriversType = 83
	JordanKing          DriversType = 84
	MahaveerRaghunathan DriversType = 85
	TatianaCalderon     DriversType = 86
	AnthoineHubert      DriversType = 87
	GuilianoAlesi       DriversType = 88
	RalphBoschung       DriversType = 89
	MichaelSchumacher   DriversType = 90
	DanTicktum          DriversType = 91
	MarcusArmstrong     DriversType = 92
	ChristianLundgaard  DriversType = 93
	YukiTsunoda         DriversType = 94
	JehanDaruvala       DriversType = 95
	GulhermeSamaia      DriversType = 96
	PedroPiquet         DriversType = 97
	FelipeDrugovich     DriversType = 98
	RoberSchwartzman    DriversType = 99
	RoyNissany          DriversType = 100
	MarinoSatorus       DriversType = 101
	AidanJackson        DriversType = 102
	CasperAkkerman      DriversType = 103
	JensonButton        DriversType = 109
	DavidCoulthard      DriversType = 110
	NicoRosberg         DriversType = 111
	OscarPiastri        DriversType = 112
	LiamLawson          DriversType = 113
	JuriVips            DriversType = 114
	TheoPourchaire      DriversType = 115
	RichardVerschoor    DriversType = 116
	LirimZendeli        DriversType = 117
	DavidBeckmann       DriversType = 118
	GianlucaPetecof     DriversType = 119
	MatteoNannini       DriversType = 120
	AlessioDeledda      DriversType = 121
	BentViscaal         DriversType = 122
	EnzoFitipaldi       DriversType = 123
)

var driverNames = map[DriversType]string{
	0:   "CarlosSainz",
	1:   "DaniilKvyat",
	2:   "DanielRicciardo",
	3:   "FernandoAlonso",
	4:   "FelipeMassa",
	6:   "KimiRaikkonen",
	7:   "LewisHamilton",
	9:   "MaxVerstappen",
	10:  "NicoHulkenburg",
	11:  "KevinMagnussen",
	12:  "RomainGrosjean",
	13:  "SebastianVettel",
	14:  "SergioPerez",
	15:  "ValtteriBottas",
	17:  "EstebanOcon",
	19:  "LanceStroll",
	20:  "ArronBarnes",
	21:  "MartinGiles",
	22:  "AlexMurray",
	23:  "LucasRoth",
	24:  "IgorCorreia",
	25:  "SophieLevasseur",
	26:  "JonasSchiffer",
	27:  "AlainForest",
	28:  "JayLetourneau",
	29:  "EstoSaari",
	30:  "YasarAtiyeh",
	31:  "CallistoCalabresi",
	32:  "NaotaIzum",
	33:  "HowardClarke",
	34:  "WilheimKaufmann",
	35:  "MarieLaursen",
	36:  "FlavioNieves",
	37:  "PeterBelousov",
	38:  "KlimekMichalski",
	39:  "SantiagoMoreno",
	40:  "BenjaminCoppens",
	41:  "NoahVisser",
	42:  "GertWaldmuller",
	43:  "JulianQuesada",
	44:  "DanielJones",
	45:  "ArtemMarkelov",
	46:  "TadasukeMakino",
	47:  "SeanGelael",
	48:  "NyckDeVries",
	49:  "JackAitken",
	50:  "GeorgeRussell",
	51:  "MaximilianGunther",
	52:  "NireiFukuzumi",
	53:  "LucaGhiotto",
	54:  "LandoNorris",
	55:  "SergioSetteCamara",
	56:  "LouisDeletraz",
	57:  "AntonioFuoco",
	58:  "CharlesLeclerc",
	59:  "PierreGasly",
	62:  "AlexanderAlbon",
	63:  "NicholasLatifi",
	64:  "DorianBoccolacci",
	65:  "NikoKari",
	66:  "RobertoMerhi",
	67:  "ArjunMaini",
	68:  "AlessioLorandi",
	69:  "RubenMeijer",
	70:  "RashidNair",
	71:  "JackTremblay",
	74:  "AntonioGiovinazzi",
	75:  "RobertKubica",
	78:  "NobuharuMatsushita",
	79:  "NikitaMazepin",
	80:  "GuanyaZhou",
	81:  "MickSchumacher",
	82:  "CallumIlott",
	83:  "JuanManuelCorrea",
	84:  "JordanKing",
	85:  "MahaveerRaghunathan",
	86:  "TatianaCalderon",
	87:  "AnthoineHubert",
	88:  "GuilianoAlesi",
	89:  "RalphBoschung",
	90:  "MichaelSchumacher",
	91:  "DanTicktum",
	92:  "MarcusArmstrong",
	93:  "ChristianLundgaard",
	94:  "YukiTsunoda",
	95:  "JehanDaruvala",
	96:  "GulhermeSamaia",
	97:  "PedroPiquet",
	98:  "FelipeDrugovich",
	99:  "RoberSchwartzman",
	100: "RoyNissany",
	101: "MarinoSatorus",
	102: "AidanJackson",
	103: "CasperAkkerman",
	109: "JensonButton",
	110: "DavidCoulthard",
	111: "NicoRosberg",
	112: "OscarPiastri",
	113: "LiamLawson",
	114: "JuriVips",
	115: "TheoPourchaire",
	116: "RichardVerschoor",
	117: "LirimZendeli",
	118: "DavidBeckmann",
	119: "GianlucaPetecof",
	120: "MatteoNannini",
	121: "AlessioDeledda",
	122: "BentViscaal",
	123: "EnzoFitipaldi",
}

func (s DriversType) String() string {
	return driverNames[s]
}
