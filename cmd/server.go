package main

import (
	"fmt"
	"strconv"

	"github.com/alexpts/go-next/next"
	"github.com/valyala/fasthttp"

	"next-project/internal/app/config"
	"next-project/internal/app/router"
)

func main() {
	conf := config.Load("./config.yml")

	app := next.NewApp()
	for _, layer := range router.FactoryLayers(conf.Layers) {
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

	fmt.Println("Starting http server on :3000")
	_ = server.ListenAndServe(conf.Http.Host + ":" + strconv.Itoa(conf.Http.Port))
}
