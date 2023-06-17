package service

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/anukul/xmon/backend/common"
	"github.com/anukul/xmon/backend/db"
	"github.com/anukul/xmon/backend/monitor"
)

type Service struct {
	logger  *zerolog.Logger
	db      common.DatabaseClient
	monitor *monitor.Monitor
}

func NewService(db *db.Database, monitor *monitor.Monitor) *Service {
	logger := log.With().Str("component", "service").Logger()
	return &Service{
		logger:  &logger,
		db:      db,
		monitor: monitor,
	}
}
