package httpcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/igordevopslabs/bjjgame/config"
)

type Server struct {
	Config *config.Config
	Router *gin.Engine
}

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})

}

func NewServer(cfg *config.Config) *Server {
	r := gin.Default()

	r.GET("/ping", Pong)

	return &Server{
		Config: cfg,
		Router: r,
	}
}
