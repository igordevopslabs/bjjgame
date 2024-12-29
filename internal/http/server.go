package httpcontroller

import (
	"fmt"

	logger "github.com/igordevopslabs/bjjgame/pkg"
	"go.uber.org/zap"
)

func (s *Server) Start() {
	port := s.Config.HTTP.Port

	if port == "" {
		port = "8080"
	}

	logger.LogInfo("Iniciando o servidor", zap.String("Port:", port))
	if err := s.Router.Run(fmt.Sprintf(":%s", port)); err != nil {
		logger.LogError("Erro ao iniciar o servidor: %v", err)
	}
}
