package view

import (
	"errors"
	"github.com/sirupsen/logrus"
	"goonbot/internal/goonbot/rtd/domain"
)

func operationToViewStr(op domain.Operation) (string, error) {
	switch op {
	case domain.Addition:
		return domain.AdditionStr, nil
	case domain.Subtraction:
		return domain.SubtractionStr, nil
	}

	logrus.Errorf("unexpected operation: %v", op)

	return "", errors.New("unexpected operation")
}
