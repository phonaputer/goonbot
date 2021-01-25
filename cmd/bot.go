package main

import (
	"errors"
	"flag"
	"github.com/sirupsen/logrus"
	"goonbot/internal/goonbot"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf, err := getConfig()
	if err != nil {
		logrus.Fatalf("error getting config: %e", err)
	}

	logrus.Info("running goonbot...")

	goonbot.Run(conf)

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	<-sigChan

	logrus.Info("shutting down goonbot")
}

func getConfig() (goonbot.Config, error) {
	discToken := flag.String("discordToken", "", "access token for Discord")
	flag.Parse()

	if discToken == nil || len(*discToken) < 1 {
		return goonbot.Config{}, errors.New("invalid config")
	}

	return goonbot.Config{
		BotToken: *discToken,
	}, nil
}
