package httpcontroller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/igordevopslabs/bjjgame/config"
	"github.com/igordevopslabs/bjjgame/pkg"
	"go.uber.org/zap"
)

type Server struct {
	Config *config.Config
	Router *gin.Engine
}

func NewServer(cfg *config.Config) *Server {
	r := gin.Default()

	r.GET("/health", Health)
	r.GET("/ready", Ready)

	return &Server{
		Config: cfg,
		Router: r,
	}
}

func (s *Server) Start() {
	port := s.Config.HTTP.Port

	if port == "" {
		port = "8080"
	}

	pkg.LogInfo("Iniciando o servidor", zap.String("Port:", port))
	if err := s.Router.Run(fmt.Sprintf(":%s", port)); err != nil {
		pkg.LogError("Erro ao iniciar o servidor: %v", err)
	}
}
