package pkg

import "github.com/igordevopslabs/bjjgame/internal/models"

func MigrateDB() {
	DB.AutoMigrate(&models.Fighters{})
	DB.AutoMigrate(&models.Techniques{})
}
