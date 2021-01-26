package handler

import (
	"goonbot/internal/goonbot/flags"
	"goonbot/internal/goonbot/subcommand/rtd"
)

const (
	purple = 6684927
)

func Rtd() CmdHandler {
	return flagErrWrapper(func(f flags.GlobalFlags, args []string) (interface{}, error) {
		res, err := rtd.RollTheDice(rtd.Settings{Language: f.Language}, args)
		if err != nil {
			return nil, err
		}

		return toDiscordEmbed("RTD", res.Title, toCodeBlock(res.Body), purple), nil
	})
}
