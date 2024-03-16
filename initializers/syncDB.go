package initializers

import (
	"jwt-token/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.Post{})
	// db.AutoMigrate(&models.Comment{})
}
