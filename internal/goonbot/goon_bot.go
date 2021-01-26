package goonbot

import (
	"context"
	"github.com/andersfylling/disgord"
	"github.com/sirupsen/logrus"
	"goonbot/internal/goonbot/rtd"
	"strings"
	"time"
)

const goonBotCommandPrefix = '!'

type Config struct {
	BotToken string
}

func Run(conf Config) {
	client := disgord.New(disgord.Config{
		BotToken: conf.BotToken,
		Logger:   logrus.New(),
		RejectEvents: []string{
			disgord.EvtTypingStart,
			disgord.EvtPresenceUpdate,
			disgord.EvtGuildMemberAdd,
			disgord.EvtGuildMemberUpdate,
			disgord.EvtGuildMemberRemove,
		},
	})
	defer client.Gateway().StayConnectedUntilInterrupted()

	client.Gateway().MessageCreate(handleMessage)
}

func handleMessage(s disgord.Session, h *disgord.MessageCreate) {
	logrus.Debug("received message")

	if !looksLikeGoonbotCommand(h) {
		return
	}

	logrus.Info("handling message")

	res, sendMsg := executeCommand(h.Message.Content)

	if sendMsg {
		h.Message.Reply(context.Background(), s, toCodeBlock(res))
	}
}

func looksLikeGoonbotCommand(h *disgord.MessageCreate) bool {
	return h != nil &&
		h.Message != nil &&
		len(h.Message.Content) > 0 &&
		h.Message.Content[0] == goonBotCommandPrefix
}

func executeCommand(msg string) (resMsg string, sendMsg bool) {
	cmd, args := splitIntoCmdAndArgs(msg)

	switch cmd {
	case "!rtd":
		return rtd.RollTheDice(args), true
	}

	return "", false
}

func splitIntoCmdAndArgs(msg string) (string, []string) {
	splitMsg := strings.Split(msg, " ")

	if len(splitMsg) > 1 {
		return splitMsg[0], splitMsg[1:]
	}

	return splitMsg[0], nil
}

func toCodeBlock(msg string) interface{} {
	return disgord.Embed{
		Title:     "Roll the Dice",
		Timestamp: disgord.Time{time.Now()},
		Color:     6684927,
		Fields:    []*disgord.EmbedField{{Name: "Your roll", Value: "```\n" + msg + "\n```"}},
	}
}
