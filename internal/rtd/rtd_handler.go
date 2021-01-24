package rtd

import (
	"fmt"
	"goonbot/internal/errutil"
	"goonbot/internal/rtd/parser"
)

func RollTheDice(input string) string {
	tokes, err := parser.Parse(input)
	if err != nil {
		return errutil.GetUserMsg(err)
	}

	total := 0
	for _, toke := range tokes {
		total = toke.Apply(total)
	}

	return fmt.Sprintf("you rolled: %v", total)
}
