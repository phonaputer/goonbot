package parser

import (
	"fmt"
	"goonbot/internal/goonbot/localization"
	"goonbot/internal/goonbot/rtd/executer"
	"strconv"
	"strings"
)

const (
	inputSeparator = " "
	diceSeparator  = "d"
)

type unparsedToken struct {
	Operation string
	Roll      string
}

func Parse(input []string) ([]executer.Token, error) {
	rawTokens, err := splitInput(input)
	if err != nil {
		return nil, err
	}

	var result []executer.Token
	for _, raw := range rawTokens {
		token, err := parseAndValidateToken(raw)
		if err != nil {
			return nil, err
		}

		result = append(result, token)
	}

	return result, nil
}

func splitInput(input []string) ([]unparsedToken, error) {
	if len(input) < 1 {
		return nil, localization.NewWithUserMsg("empty input", localization.ErrEmptyInput)
	}

	if len(input) != 1 && len(input)%2 != 1 {
		return nil, localization.NewWithUserMsg("invalid num input tokens", localization.ErrMathSeparated)
	}

	var result []unparsedToken
	result = append(result, unparsedToken{Operation: executer.AdditionStr, Roll: input[0]})
	for i := 1; i+1 < len(input); i += 2 {
		result = append(result, unparsedToken{Operation: input[i], Roll: input[i+1]})
	}

	return result, nil
}

func parseAndValidateToken(input unparsedToken) (executer.Token, error) {
	if looksLikeDiceRoll(input.Roll) {
		return parseAndValidateDiceToken(input.Operation, input.Roll)
	}

	return parseAndValidateSimpleToken(input.Operation, input.Roll)
}

func looksLikeDiceRoll(maybeRoll string) bool {
	return strings.Contains(maybeRoll, diceSeparator)
}

func parseAndValidateDiceToken(opStr, rollStr string) (*executer.DiceRoll, error) {
	op, err := executer.ParseOperation(opStr)
	if err != nil {
		return nil, fmt.Errorf("error occured while parsing roll token op: %w", err)
	}

	die, faces, err := parseRoll(rollStr)
	if err != nil {
		return nil, fmt.Errorf("error occured while parsing roll token roll: %w", err)
	}

	return &executer.DiceRoll{
		NumDie:    die,
		NumFaces:  faces,
		Operation: op,
	}, nil
}

func parseRoll(rollStr string) (die int, faces int, err error) {
	splitRoll := strings.Split(rollStr, diceSeparator)
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

func parseAndValidateSimpleToken(opStr, rollStr string) (*executer.SimpleRoll, error) {
	op, err := executer.ParseOperation(opStr)
	if err != nil {
		return nil, fmt.Errorf("error occured while parsing simple token op: %w", err)
	}

	num, err := strconv.Atoi(rollStr)
	if err != nil {
		return nil, fmt.Errorf("error occured while parsing simple token num: %w",
			localization.WithUserMsg(err, localization.ErrInvalidDiceFmt))
	}

	return &executer.SimpleRoll{
		Number:    num,
		Operation: op,
	}, nil
}
