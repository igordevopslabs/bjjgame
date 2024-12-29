package main

import (
	"github.com/igordevopslabs/bjjgame/config"
	httpcontroller "github.com/igordevopslabs/bjjgame/internal/http"
	"github.com/igordevopslabs/bjjgame/pkg"
)

func init() {
	pkg.ConnectToDatabase()
	pkg.MigrateDB()
}

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		pkg.LogError("Erro ao carregar configs", err)
	}

	server := httpcontroller.NewServer(cfg)
	server.Start()

}
