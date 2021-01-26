package handler

import (
	"github.com/andersfylling/disgord"
	"time"
)

func toCodeBlock(str string) string {
	return "```\n" + str + "\n```"
}

func toDiscordEmbed(title, name, value string, color int) disgord.Embed {
	return disgord.Embed{
		Title:     title,
		Timestamp: disgord.Time{time.Now()},
		Color:     color,
		Fields:    []*disgord.EmbedField{{Name: name, Value: value}},
	}
}
