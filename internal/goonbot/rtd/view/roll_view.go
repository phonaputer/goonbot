package view

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"goonbot/internal/goonbot/localization"
	"goonbot/internal/goonbot/rtd/domain"
	"strconv"
	"strings"
)

func rollToViewStr(roll domain.RollResult) (string, error) {
	switch typed := roll.(type) {
	case domain.DiceRollResult:
		return diceRollToViewStr(typed)
	case domain.SimpleRollResult:
		return simpleRollToViewStr(typed)
	}

	logrus.Errorf("invalid roll type: %T", roll)

	return "", errors.New("unexpected type")
}

func diceRollToViewStr(roll domain.DiceRollResult) (string, error) {
	opStr, err := operationToViewStr(roll.Operation)
	if err != nil {
		return "", fmt.Errorf("dice roll to view str operation error: %w", err)
	}

	var strRolledNums []string
	for _, rolledI := range roll.RolledNumbers {
		strRolledNums = append(strRolledNums, strconv.Itoa(rolledI))
	}
	rolledNumsForView := strings.Join(strRolledNums, ", ")

	format := localization.KeyToText(localization.DiceRollResult, localization.English)

	return fmt.Sprintf(format, opStr, roll.TotalRolled, roll.NumDie, roll.NumFaces, rolledNumsForView), nil
}

func simpleRollToViewStr(roll domain.SimpleRollResult) (string, error) {
	opStr, err := operationToViewStr(roll.Operation)
	if err != nil {
		return "", fmt.Errorf("dice roll to view str operation error: %w", err)
	}

	format := localization.KeyToText(localization.SimpleRollResult, localization.English)

	return fmt.Sprintf(format, opStr, roll.Number), nil
}
