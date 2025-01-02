package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	fighterscontroller "github.com/igordevopslabs/bjjgame/internal/controller/fighters"
	techniquescontroller "github.com/igordevopslabs/bjjgame/internal/controller/techniques"
)

func NewRouter(fightersController *fighterscontroller.FightersController, techniquesController *techniquescontroller.TechniquesController) *gin.Engine {

	router := gin.Default()

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "bjjgame is health")
	})

	router.GET("/ready", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "bjjgame is ready")
	})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	//Fighters Routes
	fighterRoutes := router.Group("/api/fighters")
	{
		fighterRoutes.POST("/", fightersController.Create)
	}
	//Techniques Route
	techniqueRoutes := router.Group("/api/techniques")
	{
		techniqueRoutes.POST("/", techniquesController.Create)
	}
	return router
}
