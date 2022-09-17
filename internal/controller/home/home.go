package home

import (
	"github.com/alexpts/go-next/next"

	"next-project/internal/controller"
	"next-project/internal/infra/logger"
)

type HomeController struct {
	controller.Controller
	logger logger.ILogger
}

func NewHomeController(logger logger.ILogger) *HomeController {
	return &HomeController{
		logger: logger,
	}
}

func (c *HomeController) HomePageAction(ctx *next.HandlerCxt) {
	c.logger.Debug(ctx, "MainPageAppHandler action run "+c.GeneralMethod())
	ctx.Response.AppendBodyString(`MainPageAppHandler`)
}
