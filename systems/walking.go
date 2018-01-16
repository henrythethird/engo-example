package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"github.com/henrythethird/game/components"
)

const walkingSpeed = 1
const jumpSpeed = 1

type walkingEntity struct {
	basic    *ecs.BasicEntity
	velocity *components.VelocityComponent
}

// The Walking system
type Walking struct {
	entities map[uint64]walkingEntity
}

// NewWalking instantiates a new walking system
func NewWalking() *Walking {
	return &Walking{make(map[uint64]walkingEntity)}
}

// Add adds an entity to the system
func (w *Walking) Add(basic *ecs.BasicEntity, velocity *components.VelocityComponent) {
	w.entities[basic.ID()] = walkingEntity{basic, velocity}
}

// Remove removes an entity from the system
func (w *Walking) Remove(basic ecs.BasicEntity) {
	delete(w.entities, basic.ID())
}

// Update increments the system by a time step
func (w *Walking) Update(dt float32) {
	left := engo.Input.Button("left").Down()
	right := engo.Input.Button("right").Down()
	jump := engo.Input.Button("jump").Down()
	horMult := float32(0)
	vertMult := float32(0)

	if left {
		horMult = -1
	}

	if right {
		horMult = 1
	}

	if jump {
		vertMult = -1
	}

	for _, entity := range w.entities {
		entity.velocity.Velocity.X = horMult * walkingSpeed

		if vertMult < 0 {
			entity.velocity.Velocity.Y = vertMult * jumpSpeed
		}
	}
}
