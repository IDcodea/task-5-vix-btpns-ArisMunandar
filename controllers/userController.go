package controllers

import (
	"encoding/json"
	"io/ioutil"

	"net/http"

	"vix-btpns/app"
	"vix-btpns/models"

	"vix-btpns/app/auth"
	"vix-btpns/helpers/errorformat"
	"vix-btpns/helpers/hash"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//Function to be used for user login
func Login(c *gin.Context) {
	//Set database
	db := c.MustGet("db").(*gorm.DB)

	//Read body form
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "Error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	//Convert json to object
	user_model := models.User{}
	err = json.Unmarshal(body, &user_model)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "Error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	//Init user
	user_model.Init()
	err = user_model.Validate("login")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "Error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	//Check if user exist
	var user_login app.UserLogin

	err = db.Debug().Table("users").Select("*").Joins("LEFT JOIN photos ON photos.user_id = users.id").
		Where("users.email = ?", user_model.Email).Find(&user_login).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "User with email " + user_model.Email + " not found",
			"data":    nil,
		})
		return
	}

	//Verify password
	err = hash.CheckPasswordHash(user_login.Password, user_model.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		formattedError := errorformat.ErrorMessage(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": formattedError.Error(),
			"data":    nil,
		})
		return
	}

	//Generate token when success login
	token, err := auth.GenerateJWT(user_login.Email, user_login.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	data := app.UserData{
		ID: user_login.ID, Username: user_login.Username, Email: user_login.Email, Token: token,
		Photos: app.Photo{Title: user_login.Title, Caption: user_login.Caption, PhotoUrl: user_login.PhotoUrl},
	}

	//Return response
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"status":  "Success",
		"message": "Login successfully",
		"data":    data,
	})
}