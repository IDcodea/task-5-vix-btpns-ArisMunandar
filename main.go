package main
import (
	"os"
	"vix-btpns/database"
	"vix-btpns/models"
	"vix-btpns/router"
)

func main() {
	db := database.ConnectDB()
	db.AutoMigrate(&models.User{})

	r := router.InitRoutes(db)
	r.Run(":" + os.Getenv("PORT"))
}