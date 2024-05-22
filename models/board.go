package models

type Board interface {
	CheckTurn(counter int) Turn
}

type board struct {
	turn Turn
}

// NewBoard creates a new instance of Board
func NewBoard() Board {
	return &board{turn: White}
}

// Define a new type called Turn
type Turn int

// Declare the enum using iota
const (
	White Turn = iota
	Black
)

func (b *board) CheckTurn(counter int) Turn {
	if counter%2 == 0 {
		b.turn = White
	} else {
		b.turn = Black
	}
	return b.turn
}

// IsWhite function to check if the turn is white
func IsWhite(b Board, counter int) bool {
	return b.CheckTurn(counter) == White
}
