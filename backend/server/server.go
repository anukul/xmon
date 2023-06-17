package server

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/anukul/xmon/backend/config"
	"github.com/anukul/xmon/backend/proto/protoconnect"
	"github.com/anukul/xmon/backend/service"
)

type Server struct {
	logger  *zerolog.Logger
	address string
	mux     *http.ServeMux
}

func NewServer(config *config.Config, service *service.Service) *Server {
	logger := log.With().Str("component", "server").Logger()

	mux := http.NewServeMux()
	path, handler := protoconnect.NewMonitorHandler(service)
	mux.Handle(path, handler)
	return &Server{
		logger:  &logger,
		address: config.ServerAddress,
		mux:     mux,
	}
}
