package game

type GameGo interface {
	CheckTurn(counter int) Turn
	IsWhite() bool
	ValidTurn()
}

type gameGo struct {
	turn        Turn
	turnCounter int
}

func (g *gameGo) ValidTurn() {
	//TODO add validation
	g.turnCounter++
}

// IsWhite function to check if the turn is white
func (g *gameGo) IsWhite() bool {
	return g.CheckTurn(g.turnCounter) == White
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

func (g *gameGo) CheckTurn(counter int) Turn {
	if counter%2 == 0 {
		g.turn = White
	} else {
		g.turn = Black
	}
	return g.turn
}
