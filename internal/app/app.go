package app

import (
	"context"
	"log/slog"

	"AnoLink/internal/api"
	_ "AnoLink/internal/logger"
	"AnoLink/internal/modules/shorter"
	"AnoLink/internal/storage"
)

func Run(ctx context.Context) error {
	store, err := storage.NewStorage(ctx)
	if err != nil {
		return err
	}
	shortener := shorter.New(store)

	router := api.NewRouter(shortener)
	server := api.NewServer(":8080", router)
	server.Start()

	gracefulShutdown(ctx)
	server.Stop(ctx)
	store.Close()

	return nil
}

func gracefulShutdown(ctx context.Context) {
	<-ctx.Done()
	slog.Info("graceful shutdown")
}
