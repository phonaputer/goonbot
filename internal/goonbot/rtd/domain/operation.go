package domain

// Operation is an enumeration of arithmetic operations supported by RTD
type Operation int

const (
	Addition Operation = iota
	Subtraction

	AdditionStr    = "+"
	SubtractionStr = "-"
)
