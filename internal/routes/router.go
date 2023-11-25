package routes

import (
	p "github.com/aakash-rajur/fold-se/internal/routes/projects"
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Router(args Args) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	Db(args.Db, router)

	Esc(args.Esc, router)

	Health(router)

	p.Projects(router)

	return router
}

type Args struct {
	Db  *sqlx.DB
	Esc *es.Client
}
