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
