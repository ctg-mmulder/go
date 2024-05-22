package game

type GameGo interface {
	CheckTurn(counter int) Turn
	IsWhite(game GameGo, turn int) bool
}

type gameGo struct {
	turn Turn
}

// IsWhite function to check if the turn is white
func (g *gameGo) IsWhite(game GameGo, turn int) bool {
	return g.CheckTurn(turn) == White
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

func (g *gameGo) CheckTurn(counter int) Turn {
	if counter%2 == 0 {
		g.turn = White
	} else {
		g.turn = Black
	}
	return g.turn
}
