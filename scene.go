package main

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"github.com/henrythethird/game/components"
	"github.com/henrythethird/game/systems"
)

// The Default scene
type Default struct{}

// Preload initializes assets
func (d *Default) Preload() {
	engo.Files.Load("icon.png")
}

// Setup function
func (d *Default) Setup(world *ecs.World) {
	d.setupKeys()

	common.SetBackground(color.White)

	gravity := systems.NewGravity()
	render := &common.RenderSystem{}
	walking := systems.NewWalking()
	network := systems.NewNetwork()

	player := components.NewPlayer()

	walking.Add(&player.BasicEntity, &player.VelocityComponent)
	gravity.Add(&player.BasicEntity, &player.VelocityComponent, &player.SpaceComponent)
	render.Add(&player.BasicEntity, &player.RenderComponent, &player.SpaceComponent)
	network.Add(&player.BasicEntity, &player.VelocityComponent, &player.SpaceComponent)

	world.AddSystem(render)
	world.AddSystem(gravity)
	world.AddSystem(walking)
	world.AddSystem(network)
}

func (d *Default) setupKeys() {
	input := engo.Input

	input.RegisterButton("left", engo.A)
	input.RegisterButton("right", engo.D)
	input.RegisterButton("jump", engo.Space)
}

// Type returns the type
func (d *Default) Type() string {
	return "Default"
}
