package main

import (
	"github.com/go-go/game/graphics"
	_ "image/png"

	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(func() {
		graphics.NewGraphics().Run()
	})
}
