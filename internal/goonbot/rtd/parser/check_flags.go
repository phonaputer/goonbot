package parser

const (
	flagPrefix  = "--"
	verboseFlag = "--v"
)

type Flags struct {
	Verbose bool
}

func CheckFlags(input []string) (flags Flags, inputMinusFlags []string) {
	foundFlags := false
	lastFlagIdx := 0
	var flgs Flags

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
