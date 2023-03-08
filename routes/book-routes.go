package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nastaro/project-api/controller"
)

func ExecuteBookRequest(router *gin.Engine) {
	router.GET("/", controller.SayHi)
	router.POST("/api/projects", controller.AddProject)
	router.GET("/api/projects", controller.GetProject)
	router.PATCH("/api/projects/:pCode", controller.UpdateProject)
	router.GET("/api/projects/:pCode", controller.GetProjectByPcode)
}
