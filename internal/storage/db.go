package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Storage Стрктурв пула соединений
type Storage struct {
	DB *pgxpool.Pool
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

// NewStorage создает новый пул соединений.
func NewStorage() *Storage {
	dsn := "postgres://demo_user:demo_password@localhost:5432/demo_db"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе: %v", err)
	}

	// Проверим соединение
	if err = pool.Ping(ctx); err != nil {
		log.Fatalf("База не отвечает: %v", err)
	}

	fmt.Println("✅ Подключение к PostgreSQL успешно!")

	return &Storage{DB: pool}
}
