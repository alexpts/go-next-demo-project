package router

import (
	"strings"

	"github.com/alexpts/go-next/next"

	"next-project/internal/app/config"
	"next-project/internal/di"
)

func FactoryLayers(routes []config.LayerConfig, actionMap di.ActionMap) []*next.Layer {
	var layers []*next.Layer
	var methods []string

	for _, route := range routes {
		handlers := actionMap[route.Action] // мапим имена на обработчики

		if route.Methods != `` {
			methods = strings.Split(route.Methods, `|`) // GET|POST|PUT
		}

		layer := next.Layer{
			Handlers:     handlers,
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
