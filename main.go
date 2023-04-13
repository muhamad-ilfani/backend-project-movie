package main

import (
	"context"
	"movie-app/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app.Run(context.Background())
}
