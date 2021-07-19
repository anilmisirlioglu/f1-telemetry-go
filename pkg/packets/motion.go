package packets

import "github.com/anilmisirlioglu/f1-telemetry-go/pkg/math"

// The motion packet gives physics data for all the cars being driven.
// There is additional data for the car being driven with the goal of being able to drive a motion platform setup.

// Frequency: Rate as specified in menus
// Size: 1464 bytes
// Version: 1

type CarMotionData struct {
	WorldPositionX     float32 // World space X position
	WorldPositionY     float32 // World space Y position
	WorldPositionZ     float32 // World space Z position
	WorldVelocityX     float32 // Velocity in world space X
	WorldVelocityY     float32 // Velocity in world space Y
	WorldVelocityZ     float32 // Velocity in world space Z
	WorldForwardDirX   int16   // World space forward X direction (normalised)
	WorldForwardDirY   int16   // World space forward Y direction (normalised)
	WorldForwardDirZ   int16   // World space forward Z direction (normalised)
	WorldRightDirX     int16   // World space right X direction (normalised)
	WorldRightDirY     int16   // World space right Y direction (normalised)
	WorldRightDirZ     int16   // World space right Z direction (normalised)
	GForceLateral      float32 // Lateral G-Force component
	GForceLongitudinal float32 // Longitudinal G-Force component
	GForceVertical     float32 // Vertical G-Force component
	Yaw                float32 // Yaw angle in radians
	Pitch              float32 // Pitch angle in radians
	Roll               float32 // Roll angle in radians
}

type PacketMotionData struct {
	Header        PacketHeader      // Header
	CarMotionData [22]CarMotionData // Data for all cars on track
	// Extra player car ONLY data
	SuspensionPosition     [4]float32 // Note: All wheel arrays have the following order:
	SuspensionVelocity     [4]float32 // RL, RR, FL, FR
	SuspensionAcceleration [4]float32 // RL, RR, FL, FR
	WheelSpeed             [4]float32 // Speed of each wheel
	WheelSlip              [4]float32 // Slip ratio for each wheel
	LocalVelocityX         float32    // Velocity in local space
	LocalVelocityY         float32    // Velocity in local space
	LocalVelocityZ         float32    // Velocity in local space
	AngularVelocityX       float32    // Angular velocity x-component
	AngularVelocityY       float32    // Angular velocity y-component
	AngularVelocityZ       float32    // Angular velocity z-component
	AngularAccelerationX   float32    // Angular velocity x-component
	AngularAccelerationY   float32    // Angular velocity y-component
	AngularAccelerationZ   float32    // Angular velocity z-component
	FrontWheelsAngle       float32    // Current front wheels angle in radians
}

func (p *PacketMotionData) LocalVelocityAsVector3() *math.Vector3 {
	return math.NewVector3(p.LocalVelocityX, p.LocalVelocityY, p.LocalVelocityZ)
}

func (p *PacketMotionData) AngularVelocityAsVector3() *math.Vector3 {
	return math.NewVector3(p.AngularVelocityX, p.AngularVelocityY, p.AngularVelocityZ)
}

func (p *PacketMotionData) AngularAccelerationAsVector3() *math.Vector3 {
	return math.NewVector3(p.AngularAccelerationX, p.AngularAccelerationY, p.AngularAccelerationZ)
}

func (p *CarMotionData) WorldPositionAsVector3() *math.Vector3 {
	return math.NewVector3(p.WorldPositionX, p.WorldPositionY, p.WorldPositionZ)
}

func (p *CarMotionData) WorldVelocityAsVector3() *math.Vector3 {
	return math.NewVector3(p.WorldVelocityX, p.WorldVelocityY, p.WorldVelocityZ)
}
