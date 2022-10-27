package router

import (
	"vix-btpns/controllers"
	"vix-btpns/middlewares"

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
	router.POST("/users/register", controllers.CreateUser)
	router.PUT("/users/:userId", controllers.UpdateUser)
	router.DELETE("/users/:userId", controllers.DeleteUser)

	router.GET("/photos", controllers.GetPhoto)

	//Middlewares for photo
	authorized := router.Group("/").Use(middlewares.AuthMiddleware())
	{
		authorized.POST("/photos", controllers.CreatePhoto)
		authorized.PUT("/photos/:photoId", controllers.UpdatePhoto)
		authorized.DELETE("/photos/:photoId", controllers.DeletePhoto)
	}
return router
}