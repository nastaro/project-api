package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nastaro/project-api/database"
	"github.com/nastaro/project-api/models"
)

func AddProject(c *gin.Context) {
	project := models.Project{}
	if err := c.BindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request, Please validate your input"})
	} else {
		if err := database.DB.Unscoped().Create(&project).Error; err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		} else {
			c.JSON(http.StatusCreated, gin.H{
				"projectName": project.ProjectName,
				"dCode":       project.Dcode,
				"ownerName":   project.OwnerName,
			})
		}
	}
}

func GetProject(c *gin.Context) {
	project := []models.Project{}
	database.DB.Find(&project)
	c.JSON(http.StatusOK, &project)
}
