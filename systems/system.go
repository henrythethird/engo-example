package systems

import (
	"engo.io/ecs"
)

// System is the base system
type System struct {
	entities map[uint64]ecs.Identifier
}

// NewSystem instantiates a new instance of System
func NewSystem() *System {
	return &System{
		make(map[uint64]ecs.Identifier),
	}
}

// AddEntity adds an entity to the system
func (s *System) AddEntity(basic ecs.Identifier) {
	s.entities[basic.ID()] = basic
}

// Remove deletes the entity from the system
func (s *System) Remove(basic ecs.BasicEntity) {
	delete(s.entities, basic.ID())
}
