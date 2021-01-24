package parser

type SimpleToken struct {
	Number    int
	Operation Operation
}

func (s SimpleToken) Apply(currentTotal int) int {
	return ExecOperation(currentTotal, s.Operation, s.Number)
}
