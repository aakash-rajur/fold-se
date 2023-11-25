package routes

import (
	e "github.com/aakash-rajur/fold-se/internal/es"
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
)

func Esc(client *es.Client, router *gin.Engine) {
	router.Use(
		func(ctx *gin.Context) {
			esCtx := e.WithEs(client, ctx.Request.Context())

			ctx.Request = ctx.Request.WithContext(esCtx)

			ctx.Next()
		},
	)
}
