package projects

import (
	cp "github.com/aakash-rajur/fold-se/internal/api/create_project"
	"github.com/aakash-rajur/fold-se/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createProject(router *gin.Engine) {
	router.POST(
		"/projects",
		func(ctx *gin.Context) {
			args := cp.Args{}

			err := ctx.ShouldBind(&args)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, utils.ErrorResult[cp.Project](err))

				return
			}

			project, err := cp.CreateProject(ctx.Request.Context(), args)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, utils.ErrorResult[cp.Project](err))

				return
			}

			ctx.JSON(http.StatusOK, utils.ValueResult(project))
		},
	)
}
