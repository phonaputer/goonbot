package handler

import (
	"github.com/andersfylling/disgord"
	"goonbot/internal/goonbot/flags"
	"goonbot/internal/goonbot/localization"
)

const (
	red = 13369344
)

type flagErrHandler func(f flags.GlobalFlags, args []string) (interface{}, error)

func flagErrWrapper(next flagErrHandler) CmdHandler {
	return func(argsAndFlags []string) interface{} {
		flgs, args, err := flags.ExtractGlobalFlags(argsAndFlags)
		if err != nil {
			return getErrEmbed(err, localization.Default)
		}

		res, err := next(flgs, args)
		if err != nil {
			return getErrEmbed(err, flgs.Language)
		}

		return res
	}
}

func getErrEmbed(err error, lang localization.Language) disgord.Embed {
	title := localization.KeyToText(localization.ErrorTitle, lang)
	name := localization.KeyToText(localization.ErrorFieldName, lang)
	value := localization.ErrToText(err, lang)

	return toDiscordEmbed(title, name, toCodeBlock(value), red)
}
