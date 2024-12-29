package pkg

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {

	var err error

	db_name := "bjj_db"
	dsn := fmt.Sprintf("host=localhost user=admin password=admin dbname=%v port=5432 sslmode=disable", db_name)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		LogError("Erro ao connectar com o banco", err)
	}

	LogInfo("Connection Ok", zap.String("DB_NAME", db_name))
}
