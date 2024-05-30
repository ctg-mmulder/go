package game

import (
	models "github.com/go-go/game/models"
)

type GameGo interface {
	CheckTurn() Turn
	IsWhiteTurn() bool
	PlayTile(x float64, y float64)
	TurnCounter() int
	IntersectState(x float64, y float64) models.State
	SetCounter(turncounter int)
}

type gameGo struct {
	turn        Turn
	turnCounter int
	intersects  []models.Intersect
}

func (g *gameGo) SetCounter(turncounter int) {
	g.turnCounter = turncounter
}

func (g *gameGo) TurnCounter() int {
	return g.turnCounter
}

func (g *gameGo) IntersectState(x float64, y float64) models.State {
	intersect, err := g.findIntersect(x, y)
	if err == nil {
		return intersect.State
	}
	panic("No intersect found")
}

func (g *gameGo) PlayTile(x float64, y float64) {
	turn := g.CheckTurn()
	intersect, err := g.findIntersect(x, y)
	if err == nil {
		g.turnCounter++
		intersect.UpdateState(toState(turn))
	}
}

func toState(turn Turn) models.State {
	if turn == White {
		return models.White
	}
	return models.Black
}

func (g *gameGo) findIntersect(x float64, y float64) (*models.Intersect, error) {
	for i := range g.intersects {
		if g.intersects[i].IsXY(x, y) {
			return &g.intersects[i], nil
		}
	}
	panic("not valid play")
}

func (g *gameGo) isFreeIntersect(intersect models.Intersect) bool {
	return intersect.State == models.Free
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
