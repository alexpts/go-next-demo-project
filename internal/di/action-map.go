package di

import (
	"net/http/pprof"

	"github.com/alexpts/go-next/next"

	"next-project/internal/controller"
	"next-project/internal/controller/hello"
)

// ActionMap - хранит маппинг между строковым action и реальными обработчиками
type ActionMap map[string][]next.Handler

// m - локальынй алиас чтобы обернуть обработчик в группу обработчиков
type m = []next.Handler

func NewActionMap(container *Container) ActionMap {
	// сервисы с их графом зависимостей достаются из контейнера и используются их методы как обработчики
	var home = container.GetHomeController()

	return ActionMap{
		"otherwise": m{controller.Otherwise},
		"mainPage":  m{home.HomePageAction},
		"mainPage2": m{home.HomePageAction},
		"hello":     m{hello.HelloAction},

		// pprof
		"/debug/pprof/":         m{next.FromHttpHandlerFunc(pprof.Index)},
		"/debug/pprof/cmdline/": m{next.FromHttpHandlerFunc(pprof.Cmdline)},
		"/debug/pprof/profile/": m{next.FromHttpHandlerFunc(pprof.Profile)},
		"/debug/pprof/symbol/":  m{next.FromHttpHandlerFunc(pprof.Symbol)},
		"/debug/pprof/trace/":   m{next.FromHttpHandlerFunc(pprof.Trace)},
	}

}
