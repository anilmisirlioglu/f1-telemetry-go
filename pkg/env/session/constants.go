package session

// Session Types
type SessionType uint8

const (
	Unknown   SessionType = 0
	P1        SessionType = 1
	P2        SessionType = 2
	P3        SessionType = 3
	ShortP    SessionType = 4
	Q1        SessionType = 5
	Q2        SessionType = 6
	Q3        SessionType = 7
	ShortQ    SessionType = 8
	OSQ       SessionType = 9
	R         SessionType = 10
	R2        SessionType = 11
	R3        SessionType = 12
	TimeTrial SessionType = 13
)

var sessionNames = map[SessionType]string{
	0:  "Unknown",
	1:  "P1",
	2:  "P2",
	3:  "P3",
	4:  "ShortP",
	5:  "Q1",
	6:  "Q2",
	7:  "Q3",
	8:  "ShortQ",
	9:  "OSQ",
	10: "R",
	11: "R2",
	12: "R3",
	13: "TimeTrial",
}

func (s SessionType) String() string {
	return sessionNames[s]
}
