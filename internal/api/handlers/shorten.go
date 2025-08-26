package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"AnoLink/internal/dto"
	"AnoLink/internal/modules/shorter"
)

type ShortenHandler struct {
	shortener *shorter.Shorter
}

func NewShortenHandler(shortener *shorter.Shorter) *ShortenHandler {
	return &ShortenHandler{
		shortener: shortener,
	}
}

func (h *ShortenHandler) HandleShorten(w http.ResponseWriter, r *http.Request) error {
	// toQRCode, _ := strconv.ParseBool(r.URL.Query().Get("qrcode"))
	// toLink, _ := strconv.ParseBool(r.URL.Query().Get("link"))

	var req dto.ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request!", http.StatusBadRequest)
		return err
	}

	if req.URL == "" {
		http.Error(w, "URL is required!", http.StatusBadRequest)
		return errors.New("URL is required")
	}

	shortURL, err := h.shortener.Shorten(req.URL, r.Host)
	if err != nil {
		http.Error(w, "Failed to Shorten!", http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.ShortenResponse{ShortURL: shortURL})

	return nil
}

func (h *ShortenHandler) HandleRedirect(w http.ResponseWriter, r *http.Request) error {
	code := chi.URLParam(r, "code")
	if code == "" {
		http.Error(w, "Code is required!", http.StatusBadRequest)
		return errors.New("code is required")
	}

	originalURL, err := h.shortener.OriginalURL(code)
	if err != nil {
		http.Error(w, "Link not found!", http.StatusNotFound)
		return err
	}

	http.Redirect(w, r, originalURL, http.StatusTemporaryRedirect)

	return nil
}
