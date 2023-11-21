package middleware

import (
	"log"
	"os"

	"githu.com/alijabbar/helpers"
	"githu.com/alijabbar/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func Authenticate(ctx *fiber.Ctx) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	secretKey := os.Getenv("TOKEN_KEY")
	authToken := ctx.Cookies("authToken")

	if authToken == "" {
		return helpers.ErrorHandler(ctx, "unAuthorized access token is required", fiber.StatusUnauthorized)
	}
	claims := jwt.MapClaims{}
	_, er := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		log.Fatal(er)
	}
	id, ok := claims["_id"].(string)
	if !ok {
		return helpers.ErrorHandler(ctx, "not a valid", fiber.StatusUnauthorized)
	}

	user, res := models.FindById(id)

	if res != nil {
		return helpers.ErrorHandler(ctx, "Unable to find", fiber.StatusUnauthorized)
	}
	ctx.Locals("user", user)

	return ctx.Next()

}
