package auth

import (
	
	"time"
	"os"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("API_SECRET"))

type ClaimJWT struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(email string, username string) (tokenString string, err error) {
	expiredTime := time.Now().Add(1 * time.Hour) //initialize expiration time
	claims := &ClaimJWT{                            //initialize claims
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //initialize token
	tokenString, err = token.SignedString(jwtKey)              //generate token string
	return
}