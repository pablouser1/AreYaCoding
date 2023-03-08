package tokens

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Create(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24)).Unix()
	claims["authorized"] = true
	claims["user"] = username

	fmt.Println(os.Getenv("APP_TOKEN"))
	tokenString, err := token.SignedString([]byte(os.Getenv("APP_TOKEN")))
	if err != nil {
		return "", err
	}

	return tokenString, err
}
