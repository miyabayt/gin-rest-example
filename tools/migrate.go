package main

import (
	"github.com/miyabayt/gin-rest/db"
	"github.com/miyabayt/gin-rest/models"
)

func main() {
	// db.DB.DropTable(&models.Pet{})
	// db.DB.DropTable(&models.Owner{})
	// db.DB.DropTable(&models.Task{})

	db.DB.AutoMigrate(&models.Pet{})
	db.DB.AutoMigrate(&models.Owner{})
	db.DB.AutoMigrate(&models.Task{})
}
