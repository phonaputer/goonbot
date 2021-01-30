package goonbot

import (
	"context"
	"errors"
	"flag"
	"github.com/andersfylling/disgord"
	"github.com/sirupsen/logrus"
	"goonbot/internal/goonbot/handler"
)

type config struct {
	botToken string
}

// Run configures the Goonbot server and path matching and then runs it
func Run() {
	conf, err := getConfig()
	if err != nil {
		logrus.Fatalf("error getting config: %e", err)
	}

	client := disgord.New(disgord.Config{
		BotToken: conf.botToken,
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

	client.Gateway().MessageCreate(handleMessage(getCmdPathRouter()))
}

func getConfig() (config, error) {
	discToken := flag.String("discordToken", "", "access token for Discord")
	flag.Parse()

	if discToken == nil || len(*discToken) < 1 {
		return config{}, errors.New("invalid config")
	}

	return config{
		botToken: *discToken,
	}, nil
}

func getCmdPathRouter() handler.CmdRouter {
	r := handler.NewCmdRouter("!goonbot", handler.UnknownCommand())
	r.HandlePath("help rtd", handler.HelpRTD())
	r.HandlePath("help", handler.Help())
	r.HandlePath("rtd", handler.Rtd())

	return r
}

func handleMessage(r handler.CmdRouter) func(s disgord.Session, h *disgord.MessageCreate) {
	return func(s disgord.Session, h *disgord.MessageCreate) {
		if h == nil || h.Message == nil {
			return
		}

		logrus.Debug("handling message")

		reply := r.Route(h.Message.Content)

		if reply != nil {
			h.Message.Reply(context.Background(), s, reply)
		}
	}
}
