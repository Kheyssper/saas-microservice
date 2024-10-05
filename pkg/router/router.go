package router

import (
	"saasmicroservice/pkg/controllers"
	"saasmicroservice/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Authenticate(db)) // Middleware de autenticação

	platformCtrl := controllers.NewPlatformController(db)

	r.POST("/platforms", platformCtrl.CreatePlatform)
	r.GET("/platforms", platformCtrl.ListPlatforms)
	// Outras rotas como Run/Stop Platforms, Delete etc.

	return r
}
