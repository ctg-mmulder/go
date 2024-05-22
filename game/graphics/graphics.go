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
	Imagesize() (float64, int)
	Boardsize() int
	Bounds() int
	Run()
	NearestValidPosition(float64) float64
	ValidMousePosition(pixel.Vec) pixel.Vec
}

type graphics struct {
	imagesize int
	boardsize int
}

// NewGraphics creates a new instance of Board
func NewGraphics(imagesize int, boardsize int) Graphics {
	return &graphics{imagesize: imagesize, boardsize: boardsize}
}

func (g graphics) Run() {
	var newGame game.GameGo = game.NewGame()
	turn := 0
	win, err := newWindowGo(g)
	imgFloatSize, imgIntSize := g.Imagesize()

	pic, err := loadPicture("./game/graphics/cross.png")
	blackpic, berr := loadPicture("./game/graphics/black-tile.png")
	whitepic, werr := loadPicture("./game/graphics/white-tile.png")
	if err != nil || berr != nil || werr != nil {
		panic(err)
	}

	cross := pixel.NewSprite(pic, pixel.R(0, 0, imgFloatSize, imgFloatSize))

	//blacktile := pixel.NewSprite(blackpic, pixel.R(0, 0, g.Imagesize(), g.Imagesize()))
	var gamematrix []pixel.Matrix
	var tiles []*pixel.Sprite

	for !win.Closed() {
		win.Clear(colornames.Aliceblue)
		woodBoard, _ := loadPicture("./game/graphics/board_wood.png")
		board := pixel.NewSprite(woodBoard, pixel.R(0, 0, float64(g.Bounds()), float64(g.Bounds())))
		board.Draw(win, pixel.IM.Moved(pixel.Vec{float64(g.Bounds() / 2), float64(g.Bounds() / 2)}))

		for x := imgIntSize; x < g.Bounds(); x += imgIntSize {
			for y := imgIntSize; y < g.Bounds(); y += imgIntSize {
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
			tile := pixel.NewSprite(pic, pixel.R(0, 0, imgFloatSize, imgFloatSize))
			tiles = append(tiles, tile)
			mouse := win.MousePosition()

			gamematrix = append(gamematrix, pixel.IM.Scaled(pixel.ZV, 1).Moved(g.ValidMousePosition(mouse)))
			fmt.Println(mouse)
			fmt.Println(g.ValidMousePosition(mouse))
		}

		win.Update()
	}
}

func (g graphics) NearestValidPosition(pos float64) float64 {
	return float64(g.roundUpToNearestImageSize(int(pos)))

}

func (g graphics) roundUpToNearestImageSize(num int) int {
	remainder := num % g.imagesize
	var result int
	if remainder >= g.imagesize/2 {
		result = ((num / g.imagesize) + 1) * g.imagesize
	} else {
		result = (num / g.imagesize) * g.imagesize
	}
	if result == 0 {
		return g.imagesize
	}
	if result >= g.imagesize*g.boardsize+g.imagesize {
		return g.imagesize * g.boardsize
	}
	return result
}

func (g graphics) Imagesize() (float64, int) {
	return float64(g.imagesize), g.imagesize
}

func (g graphics) Boardsize() int {
	return g.boardsize
}

func (g graphics) Bounds() int {
	return g.imagesize + g.imagesize*g.boardsize
}

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

func newWindowGo(graphics Graphics) (*pixelgl.Window, error) {
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

func (g graphics) ValidMousePosition(vec pixel.Vec) pixel.Vec {
	xCor := g.NearestValidPosition(vec.X)
	yCor := g.NearestValidPosition(vec.Y)

	return pixel.Vec{xCor, yCor}
}
