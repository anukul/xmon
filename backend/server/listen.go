package server

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func (s *Server) Listen() {
	s.logger.Debug().Msg(fmt.Sprintf("running on %s", s.address))
	corsHandler := cors.AllowAll().Handler(s.mux)
	if err := http.ListenAndServe(
		s.address,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(corsHandler, &http2.Server{}),
	); err != nil {
		panic(err)
	}
}
