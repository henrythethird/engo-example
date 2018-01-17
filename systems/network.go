package systems

import (
	"fmt"

	"engo.io/ecs"
	"engo.io/engo/common"

	"github.com/henrythethird/game/components"
)

type networkEntity struct {
	*ecs.BasicEntity
	*components.VelocityComponent
	*common.SpaceComponent
}

// Network system
type Network struct {
	*System
}

// NewNetwork returns an instance of the network system
func NewNetwork() *Network {
	return &Network{NewSystem()}
}

// Add adds an entity to the system
func (n *Network) Add(basic *ecs.BasicEntity, space *common.SpaceComponent, velocity *components.VelocityComponent) {
	n.AddEntity(networkEntity{basic, velocity, space})
}

// Update increments
func (n *Network) Update(dt float32) {
	for _, e := range n.entities {
		entity := e.(networkEntity)

		fmt.Println(entity.Center())
		fmt.Println(entity.Velocity)
	}
}
