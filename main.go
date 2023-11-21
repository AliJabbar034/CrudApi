package main

import (
	"fmt"
	"log"
	"os"

	"githu.com/alijabbar/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")

	app := fiber.New()

	userGroup := app.Group("/users")
	taskGroup := app.Group("/task")
	routes.UserRouter(userGroup)
	routes.TaskRouter(taskGroup)

	app.Listen(fmt.Sprintf(":%s", port))
	fmt.Println("Starting server on port ", port)

}
