package main

import (
	"image/color"

	"engo.io/engo"

	"engo.io/ecs"
	"engo.io/engo/common"
)

// The Default scene
type Default struct{}

// Preload initializes assets
func (d *Default) Preload() {
	engo.Files.Load("icon.png")
}

// Setup function
func (d *Default) Setup(world *ecs.World) {
	common.SetBackground(color.White)

	gravity := NewGravity()
	render := &common.RenderSystem{}
	player := NewPlayer()

	gravity.Add(&player.BasicEntity, &player.VelocityComponent, &player.SpaceComponent)
	render.Add(&player.BasicEntity, &player.RenderComponent, &player.SpaceComponent)

	world.AddSystem(render)
	world.AddSystem(gravity)
}

// Type returns the type
func (d *Default) Type() string {
	return "Default"
}
