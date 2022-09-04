package config

import (
	"log"
	"os"
	"strings"

	"github.com/alexpts/go-next/next"
	"gopkg.in/yaml.v3"
)

type RouterConfig struct {
	Path         string
	Methods      string
	Controller   string
	Name         string
	Priority     int
	Handler      next.Handler
	Restrictions next.Restrictions
}

func CreateLayers(projectDir string) []*next.Layer {
	rawData, err := os.ReadFile(projectDir + "/internal/app/config/router.yml")

	var routes map[string]RouterConfig
	err = yaml.Unmarshal(rawData, &routes)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return factoryLayers(routes)
}

func factoryLayers(routes map[string]RouterConfig) []*next.Layer {
	var layers []*next.Layer
	for name, route := range routes {
		handler := HandlerMap[route.Controller]
		if route.Name == `` {
			route.Name = name
		}

		var methods []string
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
