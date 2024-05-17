package main

import (
	"image"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func run() {
	imagesize := 50
	boardsize := 9
	boardbounds := 50 + imagesize*boardsize

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, float64(boardbounds), float64(boardbounds)),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	pic, err := loadPicture("cross.png")
	if err != nil {
		panic(err)
	}

	cross := pixel.NewSprite(pic, pixel.R(0, 0, 50, 50))

	for !win.Closed() {

		win.Clear(colornames.Darkolivegreen)

		for x := 50; x < boardbounds; x += 50 {
			for y := 50; y < boardbounds; y += 50 {
				cross.Draw(win, pixel.IM.Moved(pixel.Vec{float64(x), float64(y)}))
			}
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
