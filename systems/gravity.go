package systems

import (
	"engo.io/ecs"
	"engo.io/engo/common"

	"github.com/henrythethird/game/components"
)

const gravitiyIncrement = 2

type gravityComponent struct {
	*ecs.BasicEntity
	*components.VelocityComponent
	*common.SpaceComponent
}

// The Gravity system simulates gravity
type Gravity struct {
	*System
}

// NewGravity instantiates a new gravity system
func NewGravity() *Gravity {
	return &Gravity{NewSystem()}
}

// Add appends an entity to the gravity system
func (g *Gravity) Add(basic *ecs.BasicEntity, velocity *components.VelocityComponent, space *common.SpaceComponent) {
	g.AddEntity(gravityComponent{basic, velocity, space})
}

// Update simulates a time step
func (g *Gravity) Update(dt float32) {
	gravInc := dt * gravitiyIncrement

	for _, e := range g.entities {
		entity := e.(gravityComponent)

		c := entity.Center()
		c.Add(entity.GetVelocity())
		entity.SetCenter(c)

		entity.Velocity.Y += gravInc
	}
}
