package driver

// Drivers IDs
type Driver uint8

const (
	CarlosSainz         Driver = 0
	DaniilKvyat         Driver = 1
	DanielRicciardo     Driver = 2
	FernandoAlonso      Driver = 3
	FelipeMassa         Driver = 4
	KimiRaikkonen       Driver = 6
	LewisHamilton       Driver = 7
	MaxVerstappen       Driver = 9
	NicoHulkenburg      Driver = 10
	KevinMagnussen      Driver = 11
	RomainGrosjean      Driver = 12
	SebastianVettel     Driver = 13
	SergioPerez         Driver = 14
	ValtteriBottas      Driver = 15
	EstebanOcon         Driver = 17
	LanceStroll         Driver = 19
	ArronBarnes         Driver = 20
	MartinGiles         Driver = 21
	AlexMurray          Driver = 22
	LucasRoth           Driver = 23
	IgorCorreia         Driver = 24
	SophieLevasseur     Driver = 25
	JonasSchiffer       Driver = 26
	AlainForest         Driver = 27
	JayLetourneau       Driver = 28
	EstoSaari           Driver = 29
	YasarAtiyeh         Driver = 30
	CallistoCalabresi   Driver = 31
	NaotaIzum           Driver = 32
	HowardClarke        Driver = 33
	WilheimKaufmann     Driver = 34
	MarieLaursen        Driver = 35
	FlavioNieves        Driver = 36
	PeterBelousov       Driver = 37
	KlimekMichalski     Driver = 38
	SantiagoMoreno      Driver = 39
	BenjaminCoppens     Driver = 40
	NoahVisser          Driver = 41
	GertWaldmuller      Driver = 42
	JulianQuesada       Driver = 43
	DanielJones         Driver = 44
	ArtemMarkelov       Driver = 45
	TadasukeMakino      Driver = 46
	SeanGelael          Driver = 47
	NyckDeVries         Driver = 48
	JackAitken          Driver = 49
	GeorgeRussell       Driver = 50
	MaximilianGunther   Driver = 51
	NireiFukuzumi       Driver = 52
	LucaGhiotto         Driver = 53
	LandoNorris         Driver = 54
	SergioSetteCamara   Driver = 55
	LouisDeletraz       Driver = 56
	AntonioFuoco        Driver = 57
	CharlesLeclerc      Driver = 58
	PierreGasly         Driver = 59
	AlexanderAlbon      Driver = 62
	NicholasLatifi      Driver = 63
	DorianBoccolacci    Driver = 64
	NikoKari            Driver = 65
	RobertoMerhi        Driver = 66
	ArjunMaini          Driver = 67
	AlessioLorandi      Driver = 68
	RubenMeijer         Driver = 69
	RashidNair          Driver = 70
	JackTremblay        Driver = 71
	AntonioGiovinazzi   Driver = 74
	RobertKubica        Driver = 75
	NobuharuMatsushita  Driver = 78
	NikitaMazepin       Driver = 79
	GuanyaZhou          Driver = 80
	MickSchumacher      Driver = 81
	CallumIlott         Driver = 82
	JuanManuelCorrea    Driver = 83
	JordanKing          Driver = 84
	MahaveerRaghunathan Driver = 85
	TatianaCalderon     Driver = 86
	AnthoineHubert      Driver = 87
	GuilianoAlesi       Driver = 88
	RalphBoschung       Driver = 89
	MichaelSchumacher   Driver = 90
	DanTicktum          Driver = 91
	MarcusArmstrong     Driver = 92
	ChristianLundgaard  Driver = 93
	YukiTsunoda         Driver = 94
	JehanDaruvala       Driver = 95
	GulhermeSamaia      Driver = 96
	PedroPiquet         Driver = 97
	FelipeDrugovich     Driver = 98
	RoberSchwartzman    Driver = 99
	RoyNissany          Driver = 100
	MarinoSatorus       Driver = 101
	AidanJackson        Driver = 102
	CasperAkkerman      Driver = 103
	JensonButton        Driver = 109
	DavidCoulthard      Driver = 110
	NicoRosberg         Driver = 111
	OscarPiastri        Driver = 112
	LiamLawson          Driver = 113
	JuriVips            Driver = 114
	TheoPourchaire      Driver = 115
	RichardVerschoor    Driver = 116
	LirimZendeli        Driver = 117
	DavidBeckmann       Driver = 118
	GianlucaPetecof     Driver = 119
	MatteoNannini       Driver = 120
	AlessioDeledda      Driver = 121
	BentViscaal         Driver = 122
	EnzoFitipaldi       Driver = 123
)

var drivers = map[Driver]string{
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

func (s Driver) String() string {
	return drivers[s]
}
