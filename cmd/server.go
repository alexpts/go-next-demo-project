package main

import (
	"github.com/valyala/fasthttp"

	"github.com/alexpts/go-next/next"

	"next-project/internal/app/config"
)

func main() {
	app := next.NewApp()
	for _, layer := range config.CreateLayers(`.`) {
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

	_ = server.ListenAndServe(":3000")
}
