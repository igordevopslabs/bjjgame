package config

import (
	"fmt"

	"github.com/igordevopslabs/bjjgame/pkg/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//definir variaveis de conexao com banco.

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbName   = "bjj_db"
)

// Definir a função para se comunicar com o banco atraves dessas variaveis
// essa função ira retornar um objeto de comunicação com o Banco
func DatabaseConnection() *gorm.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	helper.ErrorPanic(err)

	return db
}
