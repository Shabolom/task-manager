package main

import (
	"context"
	"os"

	"github.com/labstack/gommon/log"
	"go.uber.org/zap"

	api "task_manager/gen"
	"task_manager/internal/di"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// создание контекста
	ctx := context.Background()

	// читаем env файл и записываем в локальные переменные окружения (локально на машину во время сессии)
	if err := godotenv.Load("./build/local/.env"); err != nil {
		panic(err)
	}

	// создаем новый роутер (клиент http для сетевых вызовов)
	e := echo.New()

	container := di.New(ctx)

	container.Logger()
	log.Info("starting server", zap.String("port", os.Getenv("PORT")))

	handlers := container.GetHTTPHandlers()

	api.RegisterHandlers(e, handlers)

	err := e.Start(":8080")
	if err != nil {
		log.Fatal("start server error", zap.Error(err))
	}
}
