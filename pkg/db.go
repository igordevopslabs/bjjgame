package pkg

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {

	var err error

	dsn := os.Getenv("DB_CONN_STRING")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		LogError("Erro ao connectar com o banco", err)
	}

	LogInfo("Database connection Ok")
}
