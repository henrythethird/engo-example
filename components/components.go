package components

import (
	"engo.io/engo"
)

// VelocityComponent adds velocity to the entities
type VelocityComponent struct {
	Velocity engo.Point
}

// GetVelocity returns the velocity
func (v *VelocityComponent) GetVelocity() engo.Point {
	return v.Velocity
}
