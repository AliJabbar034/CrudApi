package helpers

import "github.com/gofiber/fiber/v2"

func ErrorHandler(ctx *fiber.Ctx, message string, statusCode int) error {
	return ctx.Status(statusCode).JSON(fiber.Map{
		"message": message,
	})
}
