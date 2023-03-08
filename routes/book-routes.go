package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nastaro/project-api/controller"
)

func ExecuteBookRequest(router *gin.Engine) {
	router.POST("/api/projects", controller.AddProject)
	router.GET("/api/projects", controller.GetProject)
	// router.GET("/books/:id", controller.ShowBookById)
	// router.GET("/books/:id/review", controller.ShowBookById)
	// router.POST("/books/:id/review", controller.AddReview)
}
