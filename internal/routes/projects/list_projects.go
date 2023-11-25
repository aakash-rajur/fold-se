package projects

import (
	lp "github.com/aakash-rajur/fold-se/internal/api/list_projects"
	"github.com/aakash-rajur/fold-se/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func listProjects(router *gin.Engine) {
	router.GET(
		"/projects",
		func(ctx *gin.Context) {
			offset, limit := utils.PaginationFrom(ctx)

			args := lp.Args{
				Offset: utils.PointerTo(offset),
				Limit:  utils.PointerTo(limit),
			}

			projects, err := lp.ListProject(ctx.Request.Context(), args)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, utils.ErrorResult[[]lp.ListProjectsResult](err))

				return
			}

			ctx.JSON(http.StatusOK, utils.ValueResult(projects))
		},
	)
}
