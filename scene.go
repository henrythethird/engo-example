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
	d.setupKeys()

	common.SetBackground(color.White)

	gravity := NewGravity()
	render := &common.RenderSystem{}
	walking := NewWalking()

	player := NewPlayer()

	walking.Add(&player.BasicEntity, &player.VelocityComponent)
	gravity.Add(&player.BasicEntity, &player.VelocityComponent, &player.SpaceComponent)
	render.Add(&player.BasicEntity, &player.RenderComponent, &player.SpaceComponent)

	world.AddSystem(render)
	world.AddSystem(gravity)
	world.AddSystem(walking)
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
