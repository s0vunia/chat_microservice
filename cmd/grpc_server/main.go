package main

import (
	"context"

	"github.com/s0vunia/chat_microservice/internal/app"
	"github.com/s0vunia/chat_microservice/internal/logger"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		logger.Fatal("failed to create app", zap.Error(err))
	}

	err = a.Run()
	if err != nil {
		logger.Fatal("failed to run app", zap.Error(err))
	}
}
