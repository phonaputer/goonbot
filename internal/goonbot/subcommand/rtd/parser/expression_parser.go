package parser

import (
	"goonbot/internal/goonbot/localization"
	"goonbot/internal/goonbot/subcommand/rtd/domain"
)

type unparsedToken struct {
	Operation string
	Roll      string
}

func Parse(input []string) (*domain.Expression, error) {
	rawTokens, err := splitInput(input)
	if err != nil {
		return nil, err
	}

	var resRolls []domain.Roll
	for _, raw := range rawTokens {
		token, err := parseAndValidateRoll(raw)
		if err != nil {
			return nil, err
		}

		resRolls = append(resRolls, token)
	}

	return &domain.Expression{Rolls: resRolls}, nil
}

func splitInput(input []string) ([]unparsedToken, error) {
	if len(input) < 1 {
		return nil, localization.NewWithUserMsg("empty input", localization.ErrEmptyInput)
	}

	if len(input) != 1 && len(input)%2 != 1 {
		return nil, localization.NewWithUserMsg("invalid num input tokens", localization.ErrMathSeparated)
	}

	var result []unparsedToken
	result = append(result, unparsedToken{Operation: domain.AdditionStr, Roll: input[0]})
	for i := 1; i+1 < len(input); i += 2 {
		result = append(result, unparsedToken{Operation: input[i], Roll: input[i+1]})
	}

	return result, nil
}
