package game

import (
	"github.com/go-go/game/models"
)

type GameGo interface {
	CheckTurn() Turn
	IsWhiteTurn() bool
	UpdateTurnCounter()
}

type gameGo struct {
	turn        Turn
	turnCounter int
	intersects  []models.Intersect
}

func (g *gameGo) UpdateTurnCounter() {
	//TODO add validation
	g.turnCounter++
}

// IsWhite function to check if the turn is white
func (g *gameGo) IsWhiteTurn() bool {
	return g.CheckTurn() == White
}

// NewGame creates a new instance of Board
func NewGame() GameGo {
	return &gameGo{turn: White, turnCounter: 0, intersects: createIntersects()}
}

func createIntersects() []models.Intersect {
	var intersects []models.Intersect
	for x := 1; x < 10; x++ {
		for y := 1; y < 10; y++ {
			intersect := models.Intersect{
				Location: models.Location{X: float64(x * 50), Y: float64(y * 50)},
				State:    models.Free,
			}
			intersects = append(intersects, intersect)
		}
	}
	return intersects
}

// Define a new type called Turn
type Turn int

// Declare the enum using iota
const (
	White Turn = iota
	Black
)

func (g *gameGo) CheckTurn() Turn {
	if g.turnCounter%2 == 0 {
		g.turn = White
	} else {
		g.turn = Black
	}
	return g.turn
}
