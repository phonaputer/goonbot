package rtd

import (
	"goonbot/internal/goonbot/localization"
	"goonbot/internal/goonbot/rtd/executer"
	"goonbot/internal/goonbot/rtd/parser"
	"goonbot/internal/goonbot/rtd/view"
)

func RollTheDice(input []string) string {
	f, rollInput := parser.CheckFlags(input)

	exp, err := parser.Parse(rollInput)
	if err != nil {
		return localization.ErrToText(err, localization.English)
	}

	expRes, err := executer.ExecuteExpression(*exp)
	if err != nil {
		return localization.ErrToText(err, localization.English)
	}

	res, err := view.ExpressionToView(f, expRes)
	if err != nil {
		return localization.ErrToText(err, localization.English)
	}

	return res
}
