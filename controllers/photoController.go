package controllers

import (
	
	"net/http"
	
	
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"vix-btpns/models"
	"vix-btpns/app"
	
	
)

//GFunction to get photo profile
func GetPhoto(c *gin.Context) {
	//Create list photo
	photos := []models.Photo{}

	//Set database
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Debug().Model(&models.Photo{}).Limit(100).Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": "Photo not found",
			"data":    nil,
		})
		return
	}

	//Init list photo
	if len(photos) > 0 {
		for i := range photos {
			user := models.User{}
			err := db.Model(&models.User{}).Where("id = ?", photos[i].UserID).Take(&user).Error

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "Error",
					"message": err.Error(),
					"data":    nil,
				})
				return
			}

			photos[i].Owner = app.Owner{
				ID: user.ID, Username: user.Username, Email: user.Email,
			}
		}
	}

	//Return response
	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Data retrieved successfully",
		"data":    photos,
	})
}