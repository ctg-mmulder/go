package graphics

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"github.com/go-go/game"
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
)

type Graphics interface {
	Imagesize() int
	Boardsize() int
	Bounds() int
	Run()
}

type graphics struct {
	imagesize int
	boardsize int
}

func (g graphics) Run() {
	var newGame game.GameGo = game.NewGame()
	turn := 0
	win, err := NewWindowGo(g)

	pic, err := LoadPicture("./game/graphics/cross.png")
	blackpic, berr := LoadPicture("./game/graphics/black-tile.png")
	whitepic, werr := LoadPicture("./game/graphics/white-tile.png")
	if err != nil || berr != nil || werr != nil {
		panic(err)
	}

	cross := pixel.NewSprite(pic, pixel.R(0, 0, 50, 50))

	//blacktile := pixel.NewSprite(blackpic, pixel.R(0, 0, 50, 50))
	var gamematrix []pixel.Matrix
	var tiles []*pixel.Sprite

	for !win.Closed() {

		win.Clear(colornames.Aliceblue)
		woodBoard, _ := LoadPicture("./game/graphics/board_wood.png")
		board := pixel.NewSprite(woodBoard, pixel.R(0, 0, float64(g.Bounds()), float64(g.Bounds())))
		board.Draw(win, pixel.IM.Moved(pixel.Vec{float64(g.Bounds() / 2), float64(g.Bounds() / 2)}))

		for x := 50; x < g.Bounds(); x += 50 {
			for y := 50; y < g.Bounds(); y += 50 {
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
			if newGame.IsWhite(newGame, turn) {
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

func (g graphics) Imagesize() int {
	return g.imagesize
}

func (g graphics) Boardsize() int {
	return g.boardsize
}

func (g graphics) Bounds() int {
	return 50 + g.imagesize*g.boardsize
}

// NewGraphics creates a new instance of Board
func NewGraphics() Graphics {
	return &graphics{imagesize: 50, boardsize: 9}
}

func LoadPicture(path string) (pixel.Picture, error) {
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

func NewWindowGo(graphics Graphics) (*pixelgl.Window, error) {
	cfg := windowConfigGo(graphics)
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	return win, err
}
func windowConfigGo(graphics Graphics) pixelgl.WindowConfig {
	return pixelgl.WindowConfig{
		Title:  "Let's GO!",
		Bounds: pixel.R(0, 0, float64(graphics.Bounds()), float64(graphics.Bounds())),
		VSync:  true,
	}
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
