package systems

import (
	"engo.io/ecs"
	"engo.io/engo/common"

	"github.com/henrythethird/game/components"
)

const gravitiyIncrement = 2

type gravityComponent struct {
	basic    *ecs.BasicEntity
	velocity *components.VelocityComponent
	space    *common.SpaceComponent
}

// The Gravity system simulates gravity
type Gravity struct {
	entities map[uint64]gravityComponent
}

// NewGravity instantiates a new gravity system
func NewGravity() *Gravity {
	return &Gravity{
		entities: make(map[uint64]gravityComponent),
	}
}

// Add appends an entity to the gravity system
func (g *Gravity) Add(basic *ecs.BasicEntity, velocity *components.VelocityComponent, space *common.SpaceComponent) {
	g.entities[basic.ID()] = gravityComponent{basic, velocity, space}
}

// Remove removes the entity from the system
func (g *Gravity) Remove(basic ecs.BasicEntity) {
	delete(g.entities, basic.ID())
}

// Update simulates a time step
func (g *Gravity) Update(dt float32) {
	gravInc := dt * gravitiyIncrement

	for _, entity := range g.entities {
		c := entity.space.Center()
		c.Add(entity.velocity.Velocity)
		entity.space.SetCenter(c)

		entity.velocity.Velocity.Y += gravInc
	}
}
