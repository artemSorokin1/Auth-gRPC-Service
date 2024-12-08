package jwt

import (
	"auth_service/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"time"
)

const tokenTTL = time.Hour * 24

func NewToken(user *models.User, secretKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(tokenTTL).Unix()
	claims["email"] = user.Email

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		slog.Warn("error creating token")
		return "", err
	}

	return tokenString, nil

}
