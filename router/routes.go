package router

import (
	"vix-btpns/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//Function to initialize routes
func InitRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

		//User Routes
	router.POST("/users/login", controllers.Login)
return router
}