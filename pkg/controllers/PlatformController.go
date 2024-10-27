package controllers

import (
	"net/http"
	"saasmicroservice/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PlatformController struct {
	dbConn *gorm.DB
}

func NewPlatformController(db *gorm.DB) *PlatformController {
	return &PlatformController{dbConn: db}
}

func (ctrl *PlatformController) CreatePlatform(c *gin.Context) {
	var platform models.Platform

	if err := c.ShouldBindJSON(&platform); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.dbConn.Create(&platform).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create platform"})
		return
	}

	c.JSON(http.StatusOK, platform)
}

func (ctrl *PlatformController) ListPlatforms(c *gin.Context) {
	var platforms []models.Platform
	if err := ctrl.dbConn.Find(&platforms).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list platforms"})
		return
	}
	c.JSON(http.StatusOK, platforms)
}

func (ctrl *PlatformController) RunPlatform(c *gin.Context) {
	platformID, _ := strconv.Atoi(c.Param("platform_id"))
	var platform models.Platform

	if err := ctrl.dbConn.First(&platform, platformID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Platform not found"})
		return
	}

	platform.Status = "running"
	if err := ctrl.dbConn.Save(&platform).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start platform"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Platform started successfully"})
}

func (ctrl *PlatformController) StopPlatform(c *gin.Context) {
	platformID, _ := strconv.Atoi(c.Param("platform_id"))
	var platform models.Platform

	if err := ctrl.dbConn.First(&platform, platformID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Platform not found"})
		return
	}

	platform.Status = "stopped"
	if err := ctrl.dbConn.Save(&platform).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to stop platform"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Platform stopped successfully"})
}

func (ctrl *PlatformController) DeletePlatform(c *gin.Context) {
	platformID, _ := strconv.Atoi(c.Param("platform_id"))

	if err := ctrl.dbConn.Delete(&models.Platform{}, platformID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete platform"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Platform deleted successfully"})
}
