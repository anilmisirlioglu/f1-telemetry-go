package session

// Session Types
type Session uint8

const (
	Unknown   Session = 0
	P1        Session = 1
	P2        Session = 2
	P3        Session = 3
	ShortP    Session = 4
	Q1        Session = 5
	Q2        Session = 6
	Q3        Session = 7
	ShortQ    Session = 8
	OSQ       Session = 9
	R         Session = 10
	R2        Session = 11
	R3        Session = 12
	TimeTrial Session = 13
)

var sessions = map[Session]string{
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

func (s Session) String() string {
	return sessions[s]
}
