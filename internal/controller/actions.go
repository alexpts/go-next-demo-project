package controller

import "github.com/alexpts/go-next/next"

func Otherwise(ctx *next.HandlerCxt) {
	ctx.Response.SetStatusCode(404)
	ctx.Response.AppendBodyString(`not found`)
}
