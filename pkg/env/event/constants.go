package event

// Event String Codes
const (
	SessionStarted     string = "SSTA" // Sent when the session starts
	SessionEnded       string = "SEND" // Sent when the session ends
	FastestLap         string = "FTLP" // When a driver achieves the fastest lap
	Retirement         string = "RTMT" // When a driver retires
	DRSEnabled         string = "DRSE" // Race control have enabled DRS
	DRSDisabled        string = "DRSD" // Race control have disabled DRS
	TeamMateInPit      string = "TMPT" // Your team mate has entered the pits
	ChequeredFlag      string = "CHQF" // The chequered flag has been waved
	RaceWinner         string = "RCWN" // The race winner is announced
	PenaltyIssued      string = "PENA" // A penalty has been issued – details in event
	SpeedTrapTriggered string = "SPTP" // Speed trap has been triggered by fastest speed
)
