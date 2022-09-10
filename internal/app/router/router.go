package router

import (
	"strings"

	"github.com/alexpts/go-next/next"

	"next-project/internal/app/config"
)

func FactoryLayers(routes []config.LayerConfig) []*next.Layer {
	var layers []*next.Layer
	var methods []string

	for _, route := range routes {
		handler := HandlerMap[route.Action] // мапим имена на обработчики

		if route.Methods != `` {
			methods = strings.Split(route.Methods, `|`) // GET|POST|PUT
		}

		layer := next.Layer{
			Handlers:     []next.Handler{handler},
			Priority:     route.Priority,
			Name:         route.Name,
			Path:         route.Path,
			Restrictions: route.Restrictions,
			Methods:      methods,
		}

		layers = append(layers, &layer)
	}

	return layers
}

func MainPageAppHandler2(ctx *next.HandlerCxt) {
	ctx.Response.AppendBodyString(`MainPageAppHandler2`)
}
