package packets

// The lap data packet gives details of all the cars in the session.

// Frequency: Rate as specified in menus
// Size: 1190 bytes
// Version: 1

type LapData struct {
	LastLapTime                float32 // Last lap time in seconds
	CurrentLapTime             float32 // Current time around the lap in seconds
	Sector1TimeInMS            uint16  // Sector 1 time in milliseconds
	Sector2TimeInMS            uint16  // Sector 2 time in milliseconds
	BestLapTime                float32 // Best lap time of the session in seconds
	BestLapNum                 uint8   // Lap number best time achieved on
	BestLapSector1TimeInMS     uint16  // Sector 1 time of best lap in the session in milliseconds
	BestLapSector2TimeInMS     uint16  // Sector 2 time of best lap in the session in milliseconds
	BestLapSector3TimeInMS     uint16  // Sector 3 time of best lap in the session in milliseconds
	BestOverallSector1TimeInMS uint16  // Best overall sector 1 time of the session in milliseconds
	BestOverallSector1LapNum   uint8   // Lap number best overall sector 1 time achieved on
	BestOverallSector2TimeInMS uint16  // Best overall sector 2 time of the session in milliseconds
	BestOverallSector2LapNum   uint8   // Lap number best overall sector 2 time achieved on
	BestOverallSector3TimeInMS uint16  // Best overall sector 3 time of the session in milliseconds
	BestOverallSector3LapNum   uint8   // Lap number best overall sector 3 time achieved on
	LapDistance                float32 // Distance vehicle is around current lap in metres – could,  be negative if line hasn’t been crossed yet
	TotalDistance              float32 // Total distance travelled in session in metres – could, be negative if line hasn’t been crossed yet
	SafetyCarDelta             float32 // Delta in seconds for safety car
	CarPosition                uint8   // Car race position
	CurrentLapNum              uint8   // Current lap number
	PitStatus                  uint8   // 0 = none, 1 = pitting, 2 = in pit area
	Sector                     uint8   // 0 = sector1, 1 = sector2, 2 = sector3
	CurrentLapInvalid          uint8   // Current lap invalid - 0 = valid, 1 = invalid
	Penalties                  uint8   // Accumulated time penalties in seconds to be added
	GridPosition               uint8   // Grid position the vehicle started the race in
	DriverStatus               uint8   // Status of driver - 0 = in garage, 1 = flying lap, 2 = in lap, 3 = out lap, 4 = on track
	ResultStatus               uint8   // Result status - 0 = invalid, 1 = inactive, 2 = active, 3 = finished, 4 = disqualified, 5 = not classified, 6 = retired
}

type PacketLapData struct {
	Header  PacketHeader
	LapData [22]LapData // Lap data for all cars on track
}
