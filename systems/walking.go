package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"github.com/henrythethird/game/components"
)

const walkingSpeed = 1
const jumpSpeed = 1

type walkingEntity struct {
	*ecs.BasicEntity
	*components.VelocityComponent
}

// The Walking system
type Walking struct {
	*System
}

// NewWalking instantiates a new walking system
func NewWalking() *Walking {
	return &Walking{NewSystem()}
}

// Add adds an entity to the system
func (w *Walking) Add(basic *ecs.BasicEntity, velocity *components.VelocityComponent) {
	w.AddEntity(walkingEntity{basic, velocity})
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

	for _, e := range w.entities {
		entity := e.(walkingEntity)
		entity.Velocity.X = horMult * walkingSpeed

		if vertMult < 0 {
			entity.Velocity.Y = vertMult * jumpSpeed
		}
	}
}
