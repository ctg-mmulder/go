package main

import (
	"fmt"
	"github.com/go-go/models"
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

func getValidMousePosition(vec pixel.Vec) pixel.Vec {
	xCor := getNearestValidPosition(vec.X)
	yCor := getNearestValidPosition(vec.Y)

	return pixel.Vec{xCor, yCor}
}

func getNearestValidPosition(pos float64) float64 {
	newPos := int(pos)
	return float64(RoundUpToNearest50(newPos))
}

func RoundUpToNearest50(num int) int {
	remainder := num % 50
	var result int
	if remainder >= 25 {
		result = ((num / 50) + 1) * 50
	} else {
		result = (num / 50) * 50
	}
	if result == 0 {
		return 50
	}
	if result >= 500 {
		return 450
	}
	return result
}

func run() {
	imagesize := 50
	boardsize := 9
	boardbounds := 50 + imagesize*boardsize
	// Turn tracker
	var b models.Board = models.NewBoard()
	turn := 0

	cfg := pixelgl.WindowConfig{
		Title:  "Let's GO!",
		Bounds: pixel.R(0, 0, float64(boardbounds), float64(boardbounds)),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	pic, err := loadPicture("cross.png")
	blackpic, berr := loadPicture("black-tile.png")
	whitepic, werr := loadPicture("white-tile.png")
	if err != nil || berr != nil || werr != nil {
		panic(err)
	}

	cross := pixel.NewSprite(pic, pixel.R(0, 0, 50, 50))

	//blacktile := pixel.NewSprite(blackpic, pixel.R(0, 0, 50, 50))
	var gamematrix []pixel.Matrix
	var tiles []*pixel.Sprite

	for !win.Closed() {

		win.Clear(colornames.Aliceblue)
		woodBoard, _ := loadPicture("board_wood.png")
		board := pixel.NewSprite(woodBoard, pixel.R(0, 0, float64(boardbounds), float64(boardbounds)))
		board.Draw(win, pixel.IM.Moved(pixel.Vec{float64(boardbounds / 2), float64(boardbounds / 2)}))

		for x := 50; x < boardbounds; x += 50 {
			for y := 50; y < boardbounds; y += 50 {
				cross.Draw(win, pixel.IM.Moved(pixel.Vec{float64(x), float64(y)}))
			}
		}
		for i, tile := range tiles {
			tile.Draw(win, gamematrix[i])
		}
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			// TODO
			// Check for valid turns
			turn++
			var pic pixel.Picture
			if models.IsWhite(b, turn) {
				pic = whitepic
			} else {
				pic = blackpic
			}
			tile := pixel.NewSprite(pic, pixel.R(0, 0, 50, 50))
			tiles = append(tiles, tile)
			mouse := win.MousePosition()

			gamematrix = append(gamematrix, pixel.IM.Scaled(pixel.ZV, 1).Moved(getValidMousePosition(mouse)))
			fmt.Println(mouse)
			fmt.Println(getValidMousePosition(mouse))
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
