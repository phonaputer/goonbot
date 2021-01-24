package parser

import (
	"fmt"
	"goonbot/internal/errutil"
	"strconv"
	"strings"
)

const (
	inputSeparator = " "
	diceSeparator  = "d"

	msgEmpty                = "input must not be empty"
	msgMathSeparated        = "input must be a list of +/- separated rolls"
	msgInvalidRollFormat    = "invalid dice roll format"
	msgInvalidNumberOfDie   = "dice roll xdy: x must be > 0"
	msgInvalidNumberOfFaces = "dice roll xdy: y must be > 1"
	msgUnknownArithmetic = "unknown arithmetic operation"
)

type unparsedToken struct {
	Operation string
	Roll      string
}

func Parse(input string) ([]Token, error) {
	rawTokens, err := splitInput(input)
	if err != nil {
		return nil, err
	}

	var result []Token
	for _, raw := range rawTokens {
		token, err := parseAndValidateToken(raw)
		if err != nil {
			return nil, err
		}

		result = append(result, token)
	}

	return result, nil
}

func splitInput(input string) ([]unparsedToken, error) {
	if len(input) < 1 {
		return nil, errutil.NewWithUserMsg(msgEmpty)
	}

	args := strings.Split(input, inputSeparator)
	if len(args) != 1 && len(args)%2 != 1 {
		return nil, errutil.NewWithUserMsg(msgMathSeparated)
	}

	var result []unparsedToken
	result = append(result, unparsedToken{Operation: AdditionStr, Roll: args[0]})
	for i := 1; i+1 < len(args); i += 2 {
		result = append(result, unparsedToken{Operation: args[i], Roll: args[i+1]})
	}

	return result, nil
}

func parseAndValidateToken(input unparsedToken) (Token, error) {
	if looksLikeDiceRoll(input.Roll) {
		return parseAndValidateDiceToken(input.Operation, input.Roll)
	}

	return parseAndValidateSimpleToken(input.Operation, input.Roll)
}

func looksLikeDiceRoll(maybeRoll string) bool {
	return strings.Contains(maybeRoll, diceSeparator)
}

func parseAndValidateDiceToken(opStr, rollStr string) (*DiceToken, error) {
	op, err := ParseOperation(opStr)
	if err != nil {
		return nil, fmt.Errorf("error occured while parsing roll token op: %w", err)
	}

	die, faces, err := parseRoll(rollStr)
	if err != nil {
		return nil, fmt.Errorf("error occured while parsing roll token roll: %w", err)
	}

	return &DiceToken{
		NumDie:    die,
		NumFaces:  faces,
		Operation: op,
	}, nil
}

func parseRoll(rollStr string) (die int, faces int, err error) {
	splitRoll := strings.Split(rollStr, diceSeparator)
	if len(splitRoll) != 2 {
		return 0, 0, errutil.NewWithUserMsg(msgInvalidRollFormat)
	}

	die, err = strconv.Atoi(splitRoll[0])
	if err != nil {
		return 0, 0, errutil.WithUserMsg(err, msgInvalidRollFormat)
	}
	if die < 1 {
		return 0, 0, errutil.WithUserMsg(err, msgInvalidNumberOfDie)
	}

	faces, err = strconv.Atoi(splitRoll[1])
	if err != nil {
		return 0, 0, errutil.WithUserMsg(err, msgInvalidRollFormat)
	}
	if faces < 2 {
		return 0, 0, errutil.WithUserMsg(err, msgInvalidNumberOfFaces)
	}

	return die, faces, nil
}

func parseAndValidateSimpleToken(opStr, rollStr string) (*SimpleToken, error) {
	op, err := ParseOperation(opStr)
	if err != nil {
		return nil, fmt.Errorf("error occured while parsing simple token op: %w", err)
	}

	num, err := strconv.Atoi(rollStr)
	if err != nil {
		return nil, fmt.Errorf("error occured while parsing simple token num: %w",
			errutil.WithUserMsg(err, msgInvalidRollFormat))
	}

	return &SimpleToken{
		Number:    num,
		Operation: op,
	}, nil
}
