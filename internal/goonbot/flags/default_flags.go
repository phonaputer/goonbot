package flags

import "goonbot/internal/goonbot/localization"

// the default flags are located in their own file for ease of finding and modifying them
func defaultFlags() GlobalFlags {
	return GlobalFlags{
		Language: localization.Default,
	}
}
