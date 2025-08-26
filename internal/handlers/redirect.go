package handlers

import (
	"AnoLink/internal/storage"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RedirectHandler(store *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		if code == "" {
			http.Error(w, "Code is required", http.StatusBadRequest)
			return
		}

		originalURL, err := store.GetOriginalURL(code)
		if err != nil {
			http.Error(w, "Link not found", http.StatusNotFound)
			return
		}
		http.Redirect(w, r, originalURL, http.StatusFound)
	}
}
