package rtd

import (
	"fmt"
	"goonbot/internal/goonbot/localization"
	"goonbot/internal/goonbot/rtd/executer"
	"goonbot/internal/goonbot/rtd/parser"
)

func RollTheDice(input []string) string {
	flags, rollInput := parser.CheckFlags(input)

	rolls, err := parser.Parse(rollInput)
	if err != nil {
		return localization.ErrToText(err, localization.English)
	}

	var res executer.RollingResult
	for _, roll := range rolls {
		res = roll.Apply(res)
	}

	result := fmt.Sprintf(localization.KeyToText(localization.YouRolled, localization.English), res.Sum)

	if flags.Verbose {
		result += res.Text
	}

	return result
}
