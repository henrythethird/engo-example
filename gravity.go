package main

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

const gravitiyIncrement = 2

type gravityComponent struct {
	basic    *ecs.BasicEntity
	velocity *VelocityComponent
	space    *common.SpaceComponent
}

// The GravitySystem simulates gravity
type GravitySystem struct {
	entities map[uint64]gravityComponent
}

// NewGravity instantiates a new gravity system
func NewGravity() *GravitySystem {
	return &GravitySystem{
		entities: make(map[uint64]gravityComponent),
	}
}

// Add appends an entity to the gravity system
func (g *GravitySystem) Add(basic *ecs.BasicEntity, velocity *VelocityComponent, space *common.SpaceComponent) {
	g.entities[basic.ID()] = gravityComponent{basic, velocity, space}
}

// Remove removes the entity from the system
func (g *GravitySystem) Remove(basic ecs.BasicEntity) {
	delete(g.entities, basic.ID())
}

// Update simulates a time step
func (g *GravitySystem) Update(dt float32) {
	gravInc := dt * gravitiyIncrement

	for _, entity := range g.entities {
		c := entity.space.Center()
		c.Add(entity.velocity.Velocity)
		entity.space.SetCenter(c)

		entity.velocity.Velocity.Y += gravInc
	}
}
