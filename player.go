package main

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// Player denotes the main actor
type Player struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	VelocityComponent
}

// NewPlayer instantiates a new Player resource
func NewPlayer() *Player {
	sprite, err := common.LoadedSprite("icon.png")

	if err != nil {
		log.Fatalln(err.Error())
	}

	return &Player{
		BasicEntity: ecs.NewBasic(),
		SpaceComponent: common.SpaceComponent{
			Width:    sprite.Width() * 4,
			Height:   sprite.Height() * 4,
			Position: engo.Point{X: 490, Y: 0},
		},
		RenderComponent: common.RenderComponent{
			Drawable: sprite,
			Scale:    engo.Point{X: 4, Y: 4},
		},
	}
}
