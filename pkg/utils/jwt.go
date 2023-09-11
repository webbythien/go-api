package utils

import (
	"errors"
	"fmt"
	"lido-core/v1/pkg/configs"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func removeBearerPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}

func GenerateToken(address string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"address": address,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})
	tokenString, errToken := token.SignedString([]byte(configs.JwtSecret))
	if errToken != nil {
		return "", errors.New("failed to create token")
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (string, error) {
	tokenString = removeBearerPrefix(tokenString)
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(configs.JwtSecret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		address := claims["address"]
		// _ = claims["exp"]
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return "", errors.New("token has expired")
		}
		return address.(string), nil
	}
	return "", errors.New("token invalid")
}

func GetUserAddress(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("token is required")
	}
	address, err := ParseToken(authHeader)
	if err != nil {
		return "", err
	}
	return address, nil
}

func GetUserAddressOtps(c *fiber.Ctx) string {
	address, err := GetUserAddress(c)
	if err != nil {
		return ""
	}
	return address
}
