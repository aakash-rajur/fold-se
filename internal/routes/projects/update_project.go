package projects

import (
	up "github.com/aakash-rajur/fold-se/internal/api/update_project"
	"github.com/aakash-rajur/fold-se/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func updateProject(router *gin.Engine) {
	router.PATCH(
		"/projects/:project_id",
		func(ctx *gin.Context) {
			projectIdParam := ctx.Param("project_id")

			projectId, err := strconv.Atoi(projectIdParam)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, utils.ErrorResult[up.Project](err))

				return
			}

			args := up.Args{}

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, utils.ErrorResult[up.Project](err))

				return
			}

			args.Id = utils.PointerTo(int64(projectId))

			project, err := up.UpdateProject(ctx.Request.Context(), args)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, utils.ErrorResult[up.Project](err))

				return
			}

			ctx.JSON(http.StatusOK, utils.ValueResult(project))
		},
	)
}
