package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health(router *gin.Engine) {
	router.GET(
		"/health",
		func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		},
	)
}
