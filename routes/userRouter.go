package routes

import (
	"githu.com/alijabbar/controllers"
	"githu.com/alijabbar/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(r fiber.Router) {

	r.Post("/", controllers.CreateUser)
	r.Post("/login", controllers.LogInUser)
	r.Put("/me", middleware.Authenticate, controllers.UpdateUser)
	r.Get("/me/logout", controllers.LogOutUser)
	r.Get("/me", middleware.Authenticate, controllers.GetUser)
}
