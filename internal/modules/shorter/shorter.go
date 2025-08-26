package shorter

import (
	"fmt"
	"math/rand"
	"time"

	"AnoLink/internal/storage"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type Shorter struct {
	store    *storage.Storage
	codeSize int
}

func New(store *storage.Storage) *Shorter {
	rand.Seed(time.Now().UnixNano())

	return &Shorter{
		store:    store,
		codeSize: 6,
	}
}

func (s *Shorter) Shorten(url, host string) (string, error) {
	code := generateCode(s.codeSize)

	if err := s.store.SaveLink(code, url); err != nil {
		return "", fmt.Errorf("failed to shorten URL '%s': %w", url, err)
	}

	shortURL := fmt.Sprintf("%s/r/%s", host, code)

	return shortURL, nil
}

func (s *Shorter) OriginalURL(code string) (string, error) {
	originalURL, err := s.store.GetOriginalURL(code)
	if err != nil {
		return "", fmt.Errorf("failed to get original URL: %w", err)
	}
	return originalURL, nil
}

// generateCode генерирует случайный код длиной n символов
func generateCode(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
