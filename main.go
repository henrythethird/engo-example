package main

import (
	"engo.io/engo"
)

func main() {
	opts := engo.RunOptions{
		Title:          "Game",
		Width:          1080,
		Height:         720,
		StandardInputs: true,
		HeadlessMode:   true,
	}

	engo.Run(opts, &Default{})
}
