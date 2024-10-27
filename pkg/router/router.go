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
	r.POST("/platforms/:platform_id/run", platformCtrl.RunPlatform)   // Inicia a plataforma
	r.POST("/platforms/:platform_id/stop", platformCtrl.StopPlatform) // Para a plataforma
	r.DELETE("/platforms/:platform_id", platformCtrl.DeletePlatform)  // Exclui a plataforma
	return r
}
