package main

import (
	"fmt"
	"net/http"

	"AnoLink/internal/handlers"
	"AnoLink/internal/router"
	"AnoLink/internal/storage"
)

func main() {
	fmt.Println("Запуск сервера...")

	handlers.GenerateQRCode("https://www.twitch.tv/")

	store := storage.NewStorage()
	defer store.DB.Close()

	r := router.SetupRouter(store)

	fmt.Println("Сервер запущен на :8080")
	http.ListenAndServe(":8080", r)
}
