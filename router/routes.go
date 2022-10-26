package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//Function to initialize routes
func InitRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

return router
}