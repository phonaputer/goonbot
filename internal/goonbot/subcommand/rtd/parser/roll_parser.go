package parser

import (
	"fmt"
	"goonbot/internal/goonbot/localization"
	"goonbot/internal/goonbot/subcommand/rtd/domain"
	"strconv"
	"strings"
)

const (
	diceRollSeparator = "d"
)

func parseAndValidateRoll(input unparsedToken) (domain.Roll, error) {
	if looksLikeDiceRoll(input.Roll) {
		return parseAndValidateDiceRoll(input.Operation, input.Roll)
	}

	return parseAndValidateSimpleRoll(input.Operation, input.Roll)
}

func looksLikeDiceRoll(maybeRoll string) bool {
	return strings.Contains(maybeRoll, diceRollSeparator)
}

func parseAndValidateDiceRoll(opStr, rollStr string) (domain.DiceRoll, error) {
	op, err := parseOperation(opStr)
	if err != nil {
		return domain.DiceRoll{}, fmt.Errorf("error occured while parsing roll token op: %w", err)
	}

	die, faces, err := parseRoll(rollStr)
	if err != nil {
		return domain.DiceRoll{}, fmt.Errorf("error occured while parsing roll token roll: %w", err)
	}

	return domain.DiceRoll{
		NumDie:    die,
		NumFaces:  faces,
		Operation: op,
	}, nil
}

func parseRoll(rollStr string) (die int, faces int, err error) {
	splitRoll := strings.Split(rollStr, diceRollSeparator)
	if len(splitRoll) != 2 {
		return 0, 0, localization.NewWithUserMsg("invalid dice roll fmt", localization.ErrInvalidDiceFmt)
	}

	die, err = strconv.Atoi(splitRoll[0])
	if err != nil {
		return 0, 0, localization.WithUserMsg(err, localization.ErrInvalidDiceFmt)
	}
	if die < 1 {
		return 0, 0, localization.WithUserMsg(err, localization.ErrInvalidNumDie)
	}

	faces, err = strconv.Atoi(splitRoll[1])
	if err != nil {
		return 0, 0, localization.WithUserMsg(err, localization.ErrInvalidDiceFmt)
	}
	if faces < 2 {
		return 0, 0, localization.WithUserMsg(err, localization.ErrInvalidNumFaces)
	}

	return die, faces, nil
}

func parseAndValidateSimpleRoll(opStr, rollStr string) (domain.SimpleRoll, error) {
	op, err := parseOperation(opStr)
	if err != nil {
		return domain.SimpleRoll{}, fmt.Errorf("error occured while parsing simple token op: %w", err)
	}

	num, err := strconv.Atoi(rollStr)
	if err != nil {
		return domain.SimpleRoll{}, fmt.Errorf("error occured while parsing simple token num: %w",
			localization.WithUserMsg(err, localization.ErrInvalidDiceFmt))
	}

	return domain.SimpleRoll{
		Number:    num,
		Operation: op,
	}, nil
}
