package models

// Define the Location struct
type Location struct {
	X, Y float64
}

// Define the State enum
type State int

const (
	Free State = iota
	Black
	White
)

// Define the Intersect struct that includes a Location and a State
type Intersect struct {
	Location Location
	State    State
}

func (i Intersect) IsXY(x float64, y float64) bool {
	return i.Location.LocationMatch(x, y)
}

func (i *Intersect) UpdateState(newState State) {
	i.State = newState
}

func (l *Location) LocationMatch(x float64, y float64) bool {
	return l.X == x && l.Y == y
}
