package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/igordevopslabs/bjjgame/config"
	fighterscontroller "github.com/igordevopslabs/bjjgame/internal/controller/fighters"
	techniquescontroller "github.com/igordevopslabs/bjjgame/internal/controller/techniques"
	fightersmodel "github.com/igordevopslabs/bjjgame/internal/model/fighters"
	techniquesmodel "github.com/igordevopslabs/bjjgame/internal/model/techniques"
	fightersrepo "github.com/igordevopslabs/bjjgame/internal/repository/fighters"
	techniquesrepo "github.com/igordevopslabs/bjjgame/internal/repository/techniques"
	"github.com/igordevopslabs/bjjgame/internal/router"
	fightersservice "github.com/igordevopslabs/bjjgame/internal/service/fighters"
	techniquesservice "github.com/igordevopslabs/bjjgame/internal/service/techniques"
)

func main() {
	//Criar a engine do Gin

	//Database
	db := config.DatabaseConnection()
	validate := validator.New()

	//Realiza a migration
	db.Table("fighters").AutoMigrate(&fightersmodel.Fighters{})
	db.Table("techniques").AutoMigrate(&techniquesmodel.Techniques{})

	//Inicia o repository
	fightersRepository := fightersrepo.NewFighterRepoImpl(db)
	tecniquesRepository := techniquesrepo.NewTechniquesRepoImpl(db)
	//Inicia o Service
	fightersService := fightersservice.NewFightersServiceImpl(fightersRepository, validate)
	techniquesService := techniquesservice.NewTechniquesServiceImpl(tecniquesRepository, validate)
	//Inicia o controller
	fightersController := fighterscontroller.NewFightersController(fightersService)
	techniquesController := techniquescontroller.NewTechniquesController(techniquesService)
	//Inicia o Router
	routes := router.NewRouter(fightersController, techniquesController)

	server := &http.Server{
		Addr:           ":3000",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error to initialize http server")
		panic(err)
	}
}
