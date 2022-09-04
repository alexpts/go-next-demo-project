package config

import (
	"github.com/alexpts/go-next/next"

	"next-project/internal/controller"
	"next-project/internal/controller/hello"
	"next-project/internal/controller/home"
)

var HandlerMap = map[string]next.Handler{
	"otherwise": controller.Otherwise,
	"mainPage":  home.HomePageAction,
	"mainPage2": MainPageAppHandler2,
	"hello":     hello.HelloAction,
}
