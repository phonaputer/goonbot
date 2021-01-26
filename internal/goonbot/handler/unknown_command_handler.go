package handler

import (
	"goonbot/internal/goonbot/flags"
	"goonbot/internal/goonbot/localization"
)

func UnknownCommand() CmdHandler {
	return flagErrWrapper(func(f flags.GlobalFlags, args []string) (interface{}, error) {
		title := localization.KeyToText(localization.ErrUnknownCommandTitle, f.Language)
		name := localization.KeyToText(localization.HelpAvailableCommands, f.Language)
		value := localization.KeyToText(localization.HelpCommandDetails, f.Language)

		return toDiscordEmbed(title, name, toCodeBlock(value), red), nil
	})
}
