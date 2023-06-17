package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/anukul/xmon/backend/app"
	"github.com/anukul/xmon/backend/config"
	"github.com/anukul/xmon/backend/db"
	"github.com/anukul/xmon/backend/monitor"
	"github.com/anukul/xmon/backend/server"
	"github.com/anukul/xmon/backend/service"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.TimeOnly})

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	database, err := db.NewDatabase(cfg)
	if err != nil {
		panic(err)
	}
	var m *monitor.Monitor
	if cfg.Testing {
		if m, err = monitor.NewTestMonitor(cfg, database); err != nil {
			panic(err)
		}
	} else {
		if m, err = monitor.NewMonitor(cfg, database); err != nil {
			panic(err)
		}
	}

	svc := service.NewService(database, m)
	srv := server.NewServer(cfg, svc)

	a := app.NewApp(srv, m, cfg)
	if err := a.Run(); err != nil {
		panic(err)
	}
}
