package cryptography

import (
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type userData struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(email string) string {
	err := godotenv.Load()
	log.Info("Loading env")
    if err != nil {
        log.Error("Error loading .env file")
	}
	var key = []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &userData{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Second).Unix(),
		},
	})
	result, err := token.SignedString(key)
	if err != nil {
		log.Error("Error in signing JWT Token")
	}
    return result
}