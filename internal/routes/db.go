package routes

import (
	"github.com/aakash-rajur/fold-se/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Db(db *sqlx.DB, router *gin.Engine) {
	router.Use(
		func(ctx *gin.Context) {
			dbCtx := store.WithDb(db, ctx.Request.Context())

			ctx.Request = ctx.Request.WithContext(dbCtx)

			ctx.Next()
		},
	)
}
