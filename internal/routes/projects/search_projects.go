package projects

import (
	lp "github.com/aakash-rajur/fold-se/internal/api/list_projects"
	sp "github.com/aakash-rajur/fold-se/internal/api/search_projects"
	"github.com/aakash-rajur/fold-se/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func searchProjects(router *gin.Engine) {
	router.POST(
		"/projects/search",
		func(ctx *gin.Context) {
			args := sp.Args{}

			err := ctx.ShouldBind(&args)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, utils.ErrorResult[string](err))

				return
			}

			projects, err := sp.SearchProject(ctx.Request.Context(), args)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, utils.ErrorResult[[]lp.ListProjectsResult](err))

				return
			}

			ctx.JSON(http.StatusOK, utils.ValueResult(projects))
		},
	)
}
