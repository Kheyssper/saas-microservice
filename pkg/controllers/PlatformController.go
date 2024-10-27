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

func RunPlatform(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		platformID, _ := strconv.Atoi(c.Param("platform_id"))
		var platform models.Platform

		if err := db.First(&platform, platformID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Platform not found"})
			return
		}

		platform.Status = "running"
		db.Save(&platform)

		c.JSON(http.StatusOK, gin.H{"message": "Platform started successfully"})
	}
}

func StopPlatform(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		platformID, _ := strconv.Atoi(c.Param("platform_id"))
		var platform models.Platform

		if err := db.First(&platform, platformID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Platform not found"})
			return
		}

		platform.Status = "stopped"
		db.Save(&platform)

		c.JSON(http.StatusOK, gin.H{"message": "Platform stopped successfully"})
	}
}

func DeletePlatform(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		platformID, _ := strconv.Atoi(c.Param("platform_id"))
		if err := db.Delete(&models.Platform{}, platformID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Platform not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Platform deleted successfully"})
	}
}
