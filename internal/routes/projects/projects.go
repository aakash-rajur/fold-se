package projects

import "github.com/gin-gonic/gin"

func Projects(router *gin.Engine) {
	createProject(router)

	updateProject(router)

	listProjects(router)

	searchProjects(router)
}
