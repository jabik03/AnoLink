package router

import (
	"AnoLink/internal/handlers"
	"AnoLink/internal/storage"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// SetupRouter создаёт и настраивает маршруты
func SetupRouter(store *storage.Storage) *chi.Mux {
	rChi := chi.NewRouter()

	// Мидлвары
	rChi.Use(middleware.Logger)    // Логирование
	rChi.Use(middleware.Recoverer) // Восстановление после паники

	// Домашняя страница
	rChi.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("AnoLink API работает!"))
	})

	// Эндпоинт для сокращения ссылок
	rChi.Post("/shorten", handlers.ShortenHandler(store))

	// Эндпоинт для
	rChi.Get("/r/{code}", handlers.RedirectHandler(store))

	return rChi
}
