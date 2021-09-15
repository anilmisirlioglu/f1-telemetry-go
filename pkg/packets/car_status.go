package packets

// This packet details car statuses for all the cars in the race. It includes values such as the damage readings on the car.

// Frequency: Rate as specified in menus
// Size: 1344
// Version: 1

type CarStatusData struct {
	TractionControl         uint8   // 0 (off) - 2 (high)
	AntiLockBrakes          uint8   // 0 (off) - 1 (on)
	FuelMix                 uint8   // Fuel mix - 0 = lean, 1 = standard, 2 = rich, 3 = max
	FrontBrakeBias          uint8   // Front brake bias (percentage)
	PitLimiterStatus        uint8   // Pit limiter status - 0 = off, 1 = on
	FuelInTank              float32 // Current fuel mass
	FuelCapacity            float32 // Fuel capacity
	FuelRemainingLaps       float32 // Fuel remaining in terms of laps (value on MFD)
	MaxRPM                  uint16  // Cars max RPM, point of rev limiter
	IdleRPM                 uint16  // Cars idle RPM
	MaxGears                uint8   // Maximum number of gears
	DRSAllowed              uint8   // 0 = not allowed, 1 = allowed, -1 = unknown
	DRSActivationDistance   uint16  // 0 = DRS not available, non-zero - DRS will be available in [X] metres
	ActualTyreCompound      uint8   // See for more details: docs/TYRES.md#actual-tyre-compounds
	VisualTyreCompound      uint8   // See for more details: docs/TYRES.md#visual-tyre-compounds
	TyresAgeLaps            uint8   // Age in laps of the current set of tyres
	VehicleFIAFlags         int8    // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
	ERSStoreEnergy          float32 // ERS energy store in Joules
	ERSDeployMode           uint8   // ERS deployment mode, 0 = none, 1 = medium,  2 = overtake, 3 = hotlap
	ERSHarvestedThisLapMGUK float32 // ERS energy harvested this lap by MGU-K
	ERSHarvestedThisLapMGUH float32 // ERS energy harvested this lap by MGU-H
	ERSDeployedThisLap      float32 // ERS energy deployed this lap
	NetworkPaused           uint8   // Whether the car is paused in a network game
}

type PacketCarStatusData struct {
	Header        PacketHeader
	CarStatusData [22]CarStatusData
}
