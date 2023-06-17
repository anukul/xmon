package app

import (
	"github.com/anukul/xmon/backend/config"
	"github.com/anukul/xmon/backend/monitor"
	"github.com/anukul/xmon/backend/server"
	"os"
	"os/signal"
)

type App struct {
	server          *server.Server
	monitor         *monitor.Monitor
	refreshInterval uint
	quitChannel     chan os.Signal
}

func NewApp(server *server.Server, monitor *monitor.Monitor, config *config.Config) *App {
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, os.Interrupt)
	return &App{
		server:          server,
		monitor:         monitor,
		refreshInterval: config.RefreshInterval,
		quitChannel:     quitChannel,
	}
}
