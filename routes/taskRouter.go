package routes

import (
	"githu.com/alijabbar/controllers"
	"githu.com/alijabbar/middleware"
	"github.com/gofiber/fiber/v2"
)

func TaskRouter(r fiber.Router) {

	r.Post("/", middleware.Authenticate, controllers.CreateTask)
	r.Get("/", controllers.GetAllTasks)
	r.Get("/:id", controllers.GetATask)
	r.Delete("/:id", controllers.DeleteTask)
	r.Put("/:id", middleware.Authenticate, controllers.UpdateTask)

}
