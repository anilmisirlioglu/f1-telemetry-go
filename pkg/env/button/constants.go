package button

// Button Flags
const (
	CrossOrA        uint32 = 0x0001
	TriangleOrY     uint32 = 0x0002
	CircleOrB       uint32 = 0x0004
	SquareOrX       uint32 = 0x0008
	DPadLeft        uint32 = 0x0010
	DPadRight       uint32 = 0x0020
	DPadUp          uint32 = 0x0040
	DPadDown        uint32 = 0x0080
	OptionsOrMenu   uint32 = 0x0100
	L1OrLB          uint32 = 0x0200
	R1OrRB          uint32 = 0x0400
	L2OrRB          uint32 = 0x0800
	R2OrRT          uint32 = 0x1000
	LeftStickClick  uint32 = 0x2000
	RightStickClick uint32 = 0x4000
)
