package internal

import (
	"github.com/alexpts/go-next/next"

	"next-project/internal/infra/logger"
)

type App struct {
	logger logger.ILogger
	next.MicroApp
}

func NewApp(logger logger.ILogger) *App {
	return &App{
		logger: logger,
	}
}
