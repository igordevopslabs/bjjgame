package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	//Criar a engine do Gin

	router := gin.Default()

	//definir uma rota de probe ex: helloWorld
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "bjjgame is alive")
	})
	//Iniciar o servidor

	server := &http.Server{
		Addr:           "3000",
		Handler:        router,
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
