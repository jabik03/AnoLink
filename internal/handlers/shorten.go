package handlers

import (
	"AnoLink/internal/shortener"
	"AnoLink/internal/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func ShortenHandler(store *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//toQRCode, _ := strconv.ParseBool(r.URL.Query().Get("qrcode"))
		//toLink, _ := strconv.ParseBool(r.URL.Query().Get("link"))

		var req ShortenRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if req.URL == "" {
			http.Error(w, "URL is required", http.StatusBadRequest)
			return
		}

		code := shortener.GenerateCode(6)

		if err := store.SaveLink(code, req.URL); err != nil {
			http.Error(w, "Failed to save link", http.StatusInternalServerError)
			return
		}

		shortURL := fmt.Sprintf("http://localhost:8080/r/%s", code)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ShortenResponse{ShortURL: shortURL})
	}
}
