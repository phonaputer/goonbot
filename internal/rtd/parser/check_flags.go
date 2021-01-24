package parser

import "strings"

const (
	flagPrefix  = "--"
	verboseFlag = "--v"
)

type Flags struct {
	Verbose bool
}

func CheckFlags(input string) (flags Flags, inputMinusFlags string) {
	splitInput := strings.Split(input, " ")

	foundFlags := false
	lastFlagIdx := 0
	var flgs Flags
	for i, token := range splitInput {
		if len(token) > 1 && flagPrefix != token[:2] {
			break
		}

		foundFlags = true
		lastFlagIdx = i

		if token == verboseFlag {
			flgs.Verbose = true
		}
	}

	resInput := splitInput
	if foundFlags == true {
		resInput = splitInput[lastFlagIdx+1:]
	}

	return flgs, strings.Join(resInput, " ")
}
