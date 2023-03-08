package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nastaro/project-api/database"
	"github.com/nastaro/project-api/models"
)

func UpdateProject(c *gin.Context) {
	// Get project if exist
	var project models.Project
	if err := database.DB.Where("pCode = ?", c.Param("pCode")).First(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Project not found!"})
		return
	}

	// Validate input
	var input models.Project
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	database.DB.Model(&project).Updates(input)

	c.JSON(http.StatusOK, gin.H{"updated": project})
}

func GetProjectByPcode(c *gin.Context) {
	var project models.Project

	if err := database.DB.Where("pCode = ?", c.Param("pCode")).Find(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No project found :("})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": &project})
}
