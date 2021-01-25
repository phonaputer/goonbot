package executer

import (
	"fmt"
	"strconv"
)

type SimpleRoll struct {
	Number    int
	Operation Operation
}

func (s SimpleRoll) Apply(current RollingResult) RollingResult {
	current.Text = current.Text + fmt.Sprintf("\n%s %v", GetOperationStr(s.Operation), strconv.Itoa(s.Number))
	current.Sum = ExecOperation(current.Sum, s.Operation, s.Number)

	return current
}
