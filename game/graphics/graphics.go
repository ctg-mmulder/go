package graphics

import (
	"context"
	"github.com/faiface/pixel/pixelgl"
	"github.com/go-go/game"
	"image"
	_ "image/png"
	"os"

	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	"time"
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
	game := game.NewGame()
	win := newWindowGo(g)
	var gamematrix []pixel.Matrix
	var tiles []*pixel.Sprite
	var errorMessageDisplayed bool
	var errorMessage string
	var errorDisplayTime time.Time

	for !win.Closed() {
		if errorMessageDisplayed && time.Since(errorDisplayTime) >= 3*time.Second {
			errorMessageDisplayed = false
			errorMessage = ""
		}

		win.Clear(colornames.Aliceblue)
		g.drawBoardBackGround(win)
		g.drawCrossesOnBoard(win)
		// todo dray played times based on intersects state in the game in stead of adding them on click
		g.drawPlayedTiles(tiles, win, gamematrix)

		if errorMessageDisplayed {
			g.drawErrorMessage(win, errorMessage, g.getErrorMessageMatrix(win))
		}

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			mousePosition := g.ValidMousePosition(win.MousePosition())

			err := game.PlayTile(mousePosition.X, mousePosition.Y)
			if err != nil {
				// Show error message on board
				errorMessage = "Invalid play: " + err.Error()
				errorMessageDisplayed = true
				errorDisplayTime = time.Now()
			}
			tiles = append(tiles, g.createTile(game))
			gamematrix = append(gamematrix, g.getMousePixelMatrix(g.ValidMousePosition(mousePosition)))
		}

		win.Update()
	}
}

func (g graphics) getMousePixelMatrix(mousePosition pixel.Vec) pixel.Matrix {
	return pixel.IM.Scaled(pixel.ZV, 1).Moved(mousePosition)
}

func (g graphics) drawPlayedTiles(tiles []*pixel.Sprite, win *pixelgl.Window, gamematrix []pixel.Matrix) {
	for i, tile := range tiles {
		tile.Draw(win, gamematrix[i])
	}
}

func (g graphics) createTile(game game.GameGo) *pixel.Sprite {
	return pixel.NewSprite(g.getColorForTurn(game), pixel.R(0, 0, float64(g.imagesize), float64(g.imagesize)))
}

func (g graphics) drawCrossesOnBoard(win *pixelgl.Window) {
	pic, err := loadPicture("./game/graphics/cross.png")
	if err != nil {
		panic(err)
	}
	for x := g.imagesize; x < g.Bounds(); x += g.imagesize {
		for y := g.imagesize; y < g.Bounds(); y += g.imagesize {
			pixel.NewSprite(pic, pixel.R(0, 0, float64(g.imagesize), float64(g.imagesize))).Draw(win, pixel.IM.Moved(pixel.Vec{float64(x), float64(y)}))
		}
	}
}

func (g graphics) drawBoardBackGround(win *pixelgl.Window) {
	woodBoard, _ := loadPicture("./game/graphics/board_wood.png")
	board := pixel.NewSprite(woodBoard, pixel.R(0, 0, float64(g.Bounds()), float64(g.Bounds())))
	board.Draw(win, pixel.IM.Moved(pixel.Vec{float64(g.Bounds() / 2), float64(g.Bounds() / 2)}))
}

func (g graphics) getColorForTurn(newGame game.GameGo) pixel.Picture {
	blackpic, berr := loadPicture("./game/graphics/black-tile.png")
	whitepic, werr := loadPicture("./game/graphics/white-tile.png")
	if berr != nil {
		panic(berr)
	}
	if werr != nil {
		panic(werr)
	}
	var pic pixel.Picture
	if newGame.IsWhiteTurn() {
		pic = whitepic
	} else {
		pic = blackpic
	}
	return pic
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

func newWindowGo(graphics Graphics) *pixelgl.Window {
	cfg := windowConfigGo(graphics)
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	return win
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

func (g *graphics) drawErrorMessage(win *pixelgl.Window, message string, matrix pixel.Matrix) {
	// Load a basic font and create an atlas for text rendering
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	// Create a new text object with the loaded atlas
	basicTxt := text.New(pixel.V(0, 0), basicAtlas)
	basicTxt.Color = colornames.White

	// Clear the text box
	basicTxt.Clear()

	// Write the message into the text box
	fmt.Fprintln(basicTxt, message)

	// Get the size of the text
	textBounds := basicTxt.Bounds()

	// Create a new IMDraw instance
	imd := imdraw.New(nil)
	imd.Color = colornames.Red

	// Calculate the position for the background rectangle
	rectangleMin := textBounds.Min.Sub(pixel.V(5, 5)).Add(matrix.Project(pixel.ZV))
	rectangleMax := textBounds.Max.Add(pixel.V(5, 5)).Add(matrix.Project(pixel.ZV))

	// Draw the background rectangle
	imd.Push(rectangleMin, rectangleMax)
	imd.Rectangle(0)
	imd.Draw(win)

	// Draw the text box to the window with the provided matrix
	basicTxt.Draw(win, matrix)
}

func (g graphics) getErrorMessageMatrix(win *pixelgl.Window) pixel.Matrix {
	windowCenter := win.Bounds().Center()
	return pixel.IM.Moved(windowCenter)
}
