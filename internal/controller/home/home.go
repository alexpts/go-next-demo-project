package home

import "github.com/alexpts/go-next/next"

func HomePageAction(ctx *next.HandlerCxt) {
	ctx.Response.AppendBodyString(`MainPageAppHandler`)
}
