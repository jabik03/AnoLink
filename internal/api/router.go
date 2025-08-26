package api

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"AnoLink/internal/api/handlers"
	"AnoLink/internal/modules/shorter"
)

type handlerFunc func(http.ResponseWriter, *http.Request) error

type Router struct {
	mux *chi.Mux
	sh  *handlers.ShortenHandler
}

func NewRouter(shortener *shorter.Shorter) *Router {
	return &Router{
		mux: chi.NewRouter(),
		sh:  handlers.NewShortenHandler(shortener),
	}
}

func (router *Router) RegisterRoutes() {
	router.mux.Use(middleware.Logger, middleware.Recoverer)

	router.mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("AnoLink API работает!"))
	})

	router.mux.Route("/api/v1", func(r chi.Router) {
		r.Post("/shorten", handler(router.sh.HandleShorten))
		r.Get("/r/{code}", handler(router.sh.HandleRedirect))
	})
}

func handler(h handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			handleError(w, r, err)
		}
	}
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	slog.Error("error during request", slog.String("error", err.Error()))
}
