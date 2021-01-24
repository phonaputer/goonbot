package executer

import (
	"github.com/sirupsen/logrus"
	"goonbot/internal/localization"
)

// Operation is an enumeration of arithmetic operations supported by RTD
type Operation int

const (
	Addition Operation = iota
	Subtraction

	AdditionStr    = "+"
	SubtractionStr = "-"
)

func ParseOperation(opStr string) (Operation, error) {
	switch opStr {
	case AdditionStr:
		return Addition, nil
	case SubtractionStr:
		return Subtraction, nil
	}

	return -1, localization.NewWithUserMsg("unknown arithmetic operation", localization.ErrUnknownArithmetic)
}

func ExecOperation(left int, op Operation, right int) int {
	switch op {
	case Addition:
		return left + right
	case Subtraction:
		return left - right
	}

	logrus.Errorf("attempted to exec invalid operation: left: %d, right: %d, operation: %v", left, right, op)

	return -1
}

func GetOperationStr(op Operation) string {
	switch op {
	case Addition:
		return AdditionStr
	case Subtraction:
		return SubtractionStr
	}

	return ""
}
