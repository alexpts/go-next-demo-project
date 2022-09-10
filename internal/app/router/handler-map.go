package router

import (
	"net/http/pprof"

	"github.com/alexpts/go-next/next"

	"next-project/internal/controller"
	"next-project/internal/controller/hello"
	"next-project/internal/controller/home"
)

// HandlerMap @todo добавить пример на сервис с зависимосятми и его метод есть обработчик
var HandlerMap = map[string]next.Handler{
	"otherwise": controller.Otherwise,
	"mainPage":  home.HomePageAction,
	"mainPage2": MainPageAppHandler2,
	"hello":     hello.HelloAction,

	// pprof
	"/debug/pprof/":         next.FromHttpHandlerFunc(pprof.Index),
	"/debug/pprof/cmdline/": next.FromHttpHandlerFunc(pprof.Cmdline),
	"/debug/pprof/profile/": next.FromHttpHandlerFunc(pprof.Profile),
	"/debug/pprof/symbol/":  next.FromHttpHandlerFunc(pprof.Symbol),
	"/debug/pprof/trace/":   next.FromHttpHandlerFunc(pprof.Trace),
}
