package main

import (
	"context"
	"log"
	api "task_manager/gen"
	"task_manager/internal/di"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load("./build/local/.env"); err != nil {
		log.Println("env file not loaded:", err)
	}

	e := echo.New()

	container := di.New(ctx)

	handlers := container.GetHTTPHandlers()

	api.RegisterHandlers(e, handlers)

	e.Start(":8000")
}
