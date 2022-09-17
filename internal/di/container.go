package di

import (
	"github.com/alexpts/go-next/next"

	"next-project/internal"
	"next-project/internal/controller"
	"next-project/internal/controller/home"
	"next-project/internal/infra/logger"
)

type Container struct {
	NextMicroApp *next.MicroApp
	App          *internal.App
	Logger       logger.ILogger
	services     map[string]any

	ActionMap      ActionMap
	Controller     *controller.Controller
	HomeController *home.HomeController
}

func (c *Container) GetNextMicroApp() *next.MicroApp {
	if c.NextMicroApp == nil {
		service := next.NewApp()
		c.NextMicroApp = &service
	}
	return c.NextMicroApp
}

func (c *Container) GetApp() *internal.App {
	if c.App == nil {
		service := internal.NewApp(c.GetLogger())
		c.App = service
	}
	return c.App
}

func (c *Container) GetLogger() logger.ILogger {
	if c.Logger == nil {
		service := &logger.Logger{}
		c.Logger = service
	}
	return c.Logger
}

//##### Controllers ######

func (c *Container) GetActionsMap() ActionMap {
	if c.ActionMap == nil {
		c.ActionMap = NewActionMap(c)
	}
	return c.ActionMap
}

func (c *Container) GetHomeController() *home.HomeController {
	if c.HomeController == nil {
		c.HomeController = home.NewHomeController(c.GetLogger())
	}
	return c.HomeController
}
