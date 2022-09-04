package hello

import "github.com/alexpts/go-next/next"

func HelloAction(ctx *next.HandlerCxt) {
	params := ctx.UriParams()
	ctx.Response.AppendBodyString("Hello " + params["name"])
}
