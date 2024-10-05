package controllers

import (
	"net/http"
	"saasmicroservice/pkg/models"

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
