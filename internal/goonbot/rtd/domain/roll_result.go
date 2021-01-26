package domain

type RollResult interface {
	GetRollResultType() RollResultType
	GetTotal() int
	GetOperation() Operation
}

type DiceRollResult struct {
	Operation     Operation
	NumDie        int
	NumFaces      int
	RolledNumbers []int
	TotalRolled   int
}

func (d DiceRollResult) GetTotal() int {
	return d.TotalRolled
}

func (d DiceRollResult) GetOperation() Operation {
	return d.Operation
}

func (d DiceRollResult) GetRollResultType() RollResultType {
	return RollResultTypeDice
}

type SimpleRollResult struct {
	Operation Operation
	Number    int
}

func (s SimpleRollResult) GetRollResultType() RollResultType {
	return RollResultTypeSimple
}

func (s SimpleRollResult) GetTotal() int {
	return s.Number
}

func (s SimpleRollResult) GetOperation() Operation {
	return s.Operation
}
