package controllers

import (
	"time"

	"githu.com/alijabbar/helpers"
	"githu.com/alijabbar/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx) error {
	var user models.User

	if err := ctx.BodyParser(&user); err != nil {
		helpers.ErrorHandler(ctx, "create user", fiber.StatusBadRequest)

	}
	if user.FName == "" || user.LName == "" || user.Email == "" || user.Password == "" {
		return helpers.ErrorHandler(ctx, "Provide all the required Fields", fiber.StatusBadRequest)
	}
	password, ok := helpers.HashPassword(user.Password)
	if ok != nil {
		return helpers.ErrorHandler(ctx, "Eror during password hashing", fiber.StatusBadRequest)
	}
	user.Password = password

	_id, err := models.CreateUser(user)
	if err != nil {
		return helpers.ErrorHandler(ctx, "create user failed", fiber.StatusInternalServerError)
	}

	return helpers.SendToken(_id, ctx, user)

}

func UpdateUser(ctx *fiber.Ctx) error {
	var user models.User

	authUser := ctx.Locals("user").(models.User)
	if err := ctx.BodyParser(&user); err != nil {
		return helpers.ErrorHandler(ctx, "BodyParser error", fiber.StatusNotAcceptable)
	}
	if user.FName != "" {
		authUser.FName = user.FName
	}
	if user.LName != "" {
		authUser.LName = user.LName
	}
	if user.Email != "" {
		authUser.Email = user.Email
	}

	updateCount, err := models.UserUPdate(authUser)
	if err != nil {
		return helpers.ErrorHandler(ctx, "updation Error: ", fiber.StatusNotModified)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"updates": updateCount,
		"user":    authUser,
	})
}

func GetUser(ctx *fiber.Ctx) error {

	user, ok := ctx.Locals("user").(models.User)

	if !ok {
		return helpers.ErrorHandler(ctx, "user is not Found", fiber.StatusInternalServerError)
	}

	return ctx.JSON(fiber.Map{
		"user": user,
	})

}
func LogInUser(ctx *fiber.Ctx) error {

	var user models.User

	if err := ctx.BodyParser(&user); err != nil {
		return helpers.ErrorHandler(ctx, "Error parsing user", fiber.StatusInternalServerError)
	}

	if user.Email == "" || user.Password == "" {
		return helpers.ErrorHandler(ctx, "Provide Email and Password", fiber.StatusBadRequest)
	}

	foundUser, res := models.Login(user.Email)
	if res != nil {
		return helpers.ErrorHandler(ctx, "Provide Email and Password Correct", fiber.StatusBadRequest)
	}

	err := helpers.ComparePassword(foundUser.Password, user.Password)
	if err != nil {
		return helpers.ErrorHandler(ctx, "Provide Email and Password Correct", fiber.StatusBadRequest)
	}
	id := foundUser.ID.Hex()
	return helpers.SendToken(id, ctx, foundUser)

}

func LogOutUser(ctx *fiber.Ctx) error {
	ctx.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now(),
	})
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout successful",
	})
}
