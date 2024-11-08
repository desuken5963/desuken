// Package app provides the wrangler app.
package app

import (
	"context"
	"net/http"

	"github.com/desuken/internal/handler"
	"github.com/desuken/internal/usecase"
)

// Options is options for wrangler app.
type Options struct {
	Addr string
}

// App is the wrangler app.
type App struct {
	opts   *Options
	server *http.Server
}

// NewApp allocates and returns App.
func NewApp(opts *Options) *App {
	return &App{
		opts: opts,
	}
}

// ListenAndServe starts the http server.
// nolint: funlen
func (a *App) ListenAndServe() error {

	h := handler.NewHandler(handler.Options{
		GetUserRunner:  usecase.GetUser{},
		GetOrderRunner: usecase.GetOrder{},
	})

	f := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})

	a.server = &http.Server{
		Addr:    "localhost:3000",
		Handler: f,
	}

	return a.server.ListenAndServe()
}

// Shutdown stops the http server.
func (a *App) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}
