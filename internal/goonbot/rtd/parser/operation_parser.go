package parser

import (
	"goonbot/internal/goonbot/localization"
	"goonbot/internal/goonbot/rtd/domain"
)

func parseOperation(opStr string) (domain.Operation, error) {
	switch opStr {
	case domain.AdditionStr:
		return domain.Addition, nil
	case domain.SubtractionStr:
		return domain.Subtraction, nil
	}

	return -1, localization.NewWithUserMsg("unknown arithmetic operation", localization.ErrUnknownArithmetic)
}
