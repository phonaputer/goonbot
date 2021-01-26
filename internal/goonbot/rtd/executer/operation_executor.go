package executer

import (
	"errors"
	"github.com/sirupsen/logrus"
	"goonbot/internal/goonbot/rtd/domain"
)

func execOperation(left int, op domain.Operation, right int) (int, error) {
	switch op {
	case domain.Addition:
		return left + right, nil
	case domain.Subtraction:
		return left - right, nil
	}

	logrus.Errorf("attempted to exec invalid operation: left: %d, right: %d, operation: %v", left, right, op)

	return 0, errors.New("execOperation invalid operation")
}
