package game

type GameGo interface {
	CheckTurn(counter int) Turn
}

type gameGo struct {
	turn Turn
}

// NewGame creates a new instance of Board
func NewGame() GameGo {
	return &gameGo{turn: White}
}

// Define a new type called Turn
type Turn int

// Declare the enum using iota
const (
	White Turn = iota
	Black
)

// IsWhite function to check if the turn is white
func IsWhite(g GameGo, counter int) bool {
	return g.CheckTurn(counter) == White
}

func (g *gameGo) CheckTurn(counter int) Turn {
	if counter%2 == 0 {
		g.turn = White
	} else {
		g.turn = Black
	}
	return g.turn
}
