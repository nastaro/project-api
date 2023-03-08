package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nastaro/project-api/database"
	"github.com/nastaro/project-api/models"
)

func AddProject(c *gin.Context) {
	// id := models.Identifier{}
	// Create id
	id := models.Identifier{}
	database.DB.Create(&id)

	project := models.Project{}
	project.PCode = "P" + fmt.Sprintf("%05d", id.ID)
	if err := c.BindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request, Please validate your input"})
	} else {
		if err := database.DB.Unscoped().Create(&project).Error; err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		} else {
			c.JSON(http.StatusCreated, gin.H{
				"pCode":       project.PCode,
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

func UpdateProject(c *gin.Context) {
	// Get project if exist
	var project models.Project
	if err := database.DB.Where("p_code = ?", c.Param("pCode")).First(&project).Error; err != nil {
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

	if err := database.DB.Where("p_code = ?", c.Param("pCode")).Find(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No project found :("})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": &project})
}
