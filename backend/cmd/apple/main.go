package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/desuken/internal/app"
)

// nolint: funlen
func main() {
	// 環境変数を設定
	const gracePeriod = 10 * time.Second

	appOptions := app.Options{
		// オプションの変数を設定
	}

	app := app.NewApp(&appOptions)

	go func() {
		if err := app.ListenAndServe(); err != nil {
			slog.Error("failed to listen and serve",
				slog.String("error", err.Error()),
				slog.String("address", appOptions.Addr),
			)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), gracePeriod)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown",
			slog.String("error", err.Error()),
		)
	}

	<-ctx.Done()
	slog.Info("timeout of 30 seconds.")
	slog.Info("Server exiting")
}
