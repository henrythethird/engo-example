package main

import (
	"engo.io/ecs"
	"engo.io/engo"
)

const walkingSpeed = 1
const jumpSpeed = 1

type walkingEntity struct {
	basic    *ecs.BasicEntity
	velocity *VelocityComponent
}

// The Walking system
type WalkingSystem struct {
	entities map[uint64]walkingEntity
}

// NewWalking instantiates a new walking system
func NewWalking() *WalkingSystem {
	return &WalkingSystem{make(map[uint64]walkingEntity)}
}

// Add adds an entity to the system
func (w *WalkingSystem) Add(basic *ecs.BasicEntity, velocity *VelocityComponent) {
	w.entities[basic.ID()] = walkingEntity{basic, velocity}
}

// Remove removes an entity from the system
func (w *WalkingSystem) Remove(basic ecs.BasicEntity) {
	delete(w.entities, basic.ID())
}

// Update increments the system by a time step
func (w *WalkingSystem) Update(dt float32) {
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
