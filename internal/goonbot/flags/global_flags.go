package flags

import (
	"fmt"
	"goonbot/internal/goonbot/localization"
)

const (
	flagPrefix = "--"
)

type GlobalFlags struct {
	Language localization.Language
}

// ExtractGlobalFlags searches a list of token for global flags.
// If any flag is found its value is set in the returned GlobalFlags struct and it is removed from the result []string.
// Accordingly, the result of this function is a struct representing the status of all global flags found in the
// input (with defaults set for those not found) and a slice of all input tokens which ARE NOT global flags.
func ExtractGlobalFlags(input []string) (GlobalFlags, []string, error) {
	flgs := defaultFlags()
	var result []string
	for _, token := range input {
		if len(token) < 2 || token[:2] != flagPrefix {
			result = append(result, token)

			continue
		}

		fc := newFlagChecker()
		fc.language(token, &flgs)
		if fc.err != nil {
			return flgs, nil, fmt.Errorf("error while checking global flags: %w", fc.err)
		}

		if !fc.aFlagWasMatched {
			result = append(result, token)
		}
	}

	return flgs, result, nil
}
