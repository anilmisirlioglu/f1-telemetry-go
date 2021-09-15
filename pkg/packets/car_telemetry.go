package packets

// This packet details telemetry for all the cars in the race.
// It details various values that would be recorded on the car such as speed, throttle application, DRS etc.

// Frequency: Rate as specified in menus
// Size: 1307 bytes
// Version: 1

type CarTelemetryData struct {
	Speed                   uint16     // Speed of car in kilometres per hour
	Throttle                float32    // Amount of throttle applied (0.0 to 1.0)
	Steer                   float32    // Steering (-1.0 (full lock left) to 1.0 (full lock right))
	Brake                   float32    // Amount of brake applied (0.0 to 1.0)
	Clutch                  uint8      // Amount of clutch applied (0 to 100)
	Gear                    int8       // Gear selected (1-8, N=0, R=-1)
	EngineRPM               uint16     // Engine RPM
	DRS                     uint8      // 0 = off, 1 = on
	RevLightsPercent        uint8      // Rev lights indicator (percentage)
	BrakesTemperature       [4]uint16  // Brakes temperature (celsius)
	TyresSurfaceTemperature [4]uint8   // Tyres surface temperature (celsius)
	TyresInnerTemperature   [4]uint8   // Tyres inner temperature (celsius)
	EngineTemperature       uint16     // Engine temperature (celsius)
	TyresPressure           [4]float32 // Tyres pressure (PSI)
	SurfaceType             [4]uint8   // Driving surface, see docs/TYPES.md#surface-types
}

type PacketCarTelemetryData struct {
	Header                       PacketHeader
	CarTelemetryData             [22]CarTelemetryData
	MFDPanelIndex                uint8 // Index of MFD panel open, see for more details: docs/CAR.md#mfd-index-panel
	MFDPanelIndexSecondaryPlayer uint8 // See above
	SuggestedGear                int8  // Suggested gear for the player (1-8), 0 if no gear suggested
}
