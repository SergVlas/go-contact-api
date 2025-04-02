package http

import (
	"app1/internal/config"
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handlers http.Handler) (*Server, error) {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + cfg.HTTP.Port,
			Handler: handlers,
		},
	}, nil
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
