package game

type GameGo interface {
	CheckTurn() Turn
	IsWhiteTurn() bool
	UpdateTurnCounter()
}

type gameGo struct {
	turn        Turn
	turnCounter int
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
	return &gameGo{turn: White, turnCounter: 0}
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
