package main

import (
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"

	"next-project/internal/app/config"
	"next-project/internal/app/router"
	"next-project/internal/di"
)

func main() {
	conf := config.Load("./config.yml")
	container := di.Container{}
	app := container.GetNextMicroApp()

	for _, layer := range router.FactoryLayers(conf.Layers, container.GetActionsMap()) {
		app.AddLayer(layer)
	}

	server := &fasthttp.Server{
		Handler:                       app.FasthttpHandler,
		NoDefaultDate:                 true,
		NoDefaultContentType:          true,
		NoDefaultServerHeader:         true,
		TCPKeepalive:                  true,
		GetOnly:                       true,
		DisableHeaderNamesNormalizing: true,
	}

	bindAddress := conf.Http.Host + ":" + strconv.Itoa(conf.Http.Port)
	fmt.Println(fmt.Sprintf("Starting http server on %s", bindAddress))
	_ = server.ListenAndServe(bindAddress)
}
