package helpers

import (
	"os"
	"time"

	"githu.com/alijabbar/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func SendToken(id string, c *fiber.Ctx, user models.User) error {
	token, eror := GenerateToken(id)
	if eror != nil {
		return eror
	}

	expiration := time.Now().Add(2 * 30 * 24 * time.Hour)

	c.Cookie(&fiber.Cookie{
		Name:    "authToken",
		Value:   token,
		Expires: expiration,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
		"user":  user,
	})

}

func GenerateToken(id string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	secKey := os.Getenv("TOKEN_KEY")
	expirationTime := time.Now().Add(2 * 30 * 24 * time.Hour)
	claims := jwt.MapClaims{
		"_id": id,
		"exp": expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte(secKey)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
