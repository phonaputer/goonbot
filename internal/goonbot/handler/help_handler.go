package handler

import (
	"goonbot/internal/goonbot/flags"
	"goonbot/internal/goonbot/localization"
)

func Help() CmdHandler {
	return flagErrWrapper(func(f flags.GlobalFlags, args []string) (interface{}, error) {
		title := localization.KeyToText(localization.HelpTitle, f.Language)
		name := localization.KeyToText(localization.HelpAvailableCommands, f.Language)
		value := localization.KeyToText(localization.HelpCommandDetails, f.Language)

		return toDiscordEmbed(title, name, toCodeBlock(value), purple), nil
	})
}
