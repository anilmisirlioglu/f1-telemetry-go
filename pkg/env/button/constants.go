package button

// Button Flags
const (
	CrossOrA        uint32 = 0x00000001
	TriangleOrY     uint32 = 0x00000002
	CircleOrB       uint32 = 0x00000004
	SquareOrX       uint32 = 0x00000008
	DPadLeft        uint32 = 0x00000010
	DPadRight       uint32 = 0x00000020
	DPadUp          uint32 = 0x00000040
	DPadDown        uint32 = 0x00000080
	OptionsOrMenu   uint32 = 0x00000100
	L1OrLB          uint32 = 0x00000200
	R1OrRB          uint32 = 0x00000400
	L2OrRB          uint32 = 0x00000800
	R2OrRT          uint32 = 0x00001000
	LeftStickClick  uint32 = 0x00002000
	RightStickClick uint32 = 0x00004000
	RightStickLeft  uint32 = 0x00008000
	RightStickRight uint32 = 0x00010000
	RightStickUp    uint32 = 0x00020000
)
