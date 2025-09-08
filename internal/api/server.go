package api

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(addr string, router *Router) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: router.mux,
		},
	}
}

func (s *Server) Start() {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()
	slog.Info("server started", slog.String("addr", s.httpServer.Addr))
}

func (s *Server) Stop(ctx context.Context) {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		slog.Error("terminate", slog.String("error", err.Error()))
	}
	slog.Info("server stopped")
}
