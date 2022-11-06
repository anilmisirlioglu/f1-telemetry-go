package nationality

// Nationality IDs
//
//go:generate stringer -type=Nationality -linecomment
type Nationality uint8

const (
	American Nationality = iota + 1
	Argentinean
	Australian
	Austrian
	Azerbaijani
	Bahraini
	Belgian
	Bolivian
	Brazilian
	British
	Bulgarian
	Cameroonian
	Canadian
	Chilean
	Chinese
	Colombian
	CostaRican
	Croatian
	Cypriot
	Czech
	Danish
	Dutch
	Ecuadorian
	English
	Emirian
	Estonian
	Finnish
	French
	German
	Ghanaian
	Greek
	Guatemalan
	Honduran
	HongKonger // Hong Konger
	Hungarian
	Icelander
	Indian
	Indonesian
	Irish
	Israeli
	Italian
	Jamaican
	Japanese
	Jordanian
	Kuwaiti
	Latvian
	Lebanese
	Lithuanian
	Luxembourger
	Malaysian
	Maltese
	Mexican
	Monegasque
	NewZealander // New Zealander
	Nicaraguan
	NorthernIrish // Northern Irish
	Norwegian
	Omani
	Pakistani
	Panamanian
	Paraguayan
	Peruvian
	Polish
	Portuguese
	Qatari
	Romanian
	Russian
	Salvadoran
	Saudi
	Scottish
	Serbian
	Singaporean
	Slovakian
	Slovenian
	SouthKorean  // South Korean
	SouthAfrican // South African
	Spanish
	Swedish
	Swiss
	Thai
	Turkish
	Uruguayan
	Ukrainian
	Venezuelan
	Welsh
	Barbadian
	Vietnamese
)
