package domain

type Roll interface {
	RollType() RollType
}

type SimpleRoll struct {
	Number    int
	Operation Operation
}

func (s SimpleRoll) RollType() RollType {
	return RollTypeSimple
}

type DiceRoll struct {
	NumDie    int
	NumFaces  int
	Operation Operation
}

func (d DiceRoll) RollType() RollType {
	return RollTypeDice
}
