package storage

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Storage структура пула соединений
type Storage struct {
	DB *pgxpool.Pool
}

// NewStorage создает новый пул соединений.
func NewStorage(ctx context.Context) (*Storage, error) {
	dsn := "postgres://demo_user:demo_password@localhost:5432/demo_db"

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("new pgxpool: %v", err)
	}

	// Проверим соединение
	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ping pgxpool: %v", err)
	}

	slog.Debug("connected to postgres")

	return &Storage{DB: pool}, nil
}

// Close завершает соединения.
func (s *Storage) Close() {
	s.DB.Close()
	slog.Debug("disconnected from postgres")
}

func (s *Storage) SaveLink(code, originalURL string) error {
	query := `
		INSERT INTO links (code, original_url)
		VALUES ($1, $2)
		`
	_, err := s.DB.Exec(context.Background(), query, code, originalURL)
	return err
}

func (s *Storage) GetOriginalURL(code string) (string, error) {
	query := `SELECT original_url FROM links WHERE code = $1`
	var url string
	err := s.DB.QueryRow(context.Background(), query, code).Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}
