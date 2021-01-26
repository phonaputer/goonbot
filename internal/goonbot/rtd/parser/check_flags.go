package parser

import "goonbot/internal/goonbot/rtd/domain"

const (
	flagPrefix  = "--"
	verboseFlag = "--v"
)

func CheckFlags(input []string) (flags domain.RTDFlags, inputMinusFlags []string) {
	foundFlags := false
	lastFlagIdx := 0
	var flgs domain.RTDFlags

	for i, token := range input {
		if len(token) > 1 && flagPrefix != token[:2] {
			break
		}

		foundFlags = true
		lastFlagIdx = i

		if token == verboseFlag {
			flgs.Verbose = true
		}
	}

	if foundFlags == true && len(input) > lastFlagIdx+1 {
		return flgs, input[lastFlagIdx+1:]
	}

	return flgs, input
}
