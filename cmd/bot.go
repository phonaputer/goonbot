package main

import (
	"github.com/sirupsen/logrus"
	"goonbot/internal/goonbot"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.Info("running goonbot...")

	goonbot.Run()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	<-sigChan

	logrus.Info("stopping goonbot")
}
