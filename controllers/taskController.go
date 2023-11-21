package controllers

import (
	"time"

	"githu.com/alijabbar/helpers"
	"githu.com/alijabbar/models"
	"github.com/gofiber/fiber/v2"
)

func CreateTask(ctx *fiber.Ctx) error {
	var task models.Tasks
	user, ok := ctx.Locals("user").(models.User)
	if !ok {
		return helpers.ErrorHandler(ctx, "User not found", fiber.StatusInternalServerError)
	}

	if err := ctx.BodyParser(&task); err != nil {
		return helpers.ErrorHandler(ctx, "Error parsing task", fiber.StatusBadRequest)
	}
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	task.User = user
	id, er := task.TaskCreate()
	if er != nil {
		return helpers.ErrorHandler(ctx, "  error Creating Task", fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"task": task,
		"id":   id,
	})

}

func GetAllTasks(ctx *fiber.Ctx) error {

	tasks, err := models.AllTask()
	if err != nil {
		return helpers.ErrorHandler(ctx, "GetAllTasks  failed", fiber.StatusNotFound)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"task": tasks,
	})
}

func GetATask(ctx *fiber.Ctx) error {
	var task models.Tasks

	_id := ctx.Params("id")
	err := task.GetATasks(_id)
	if err != nil {
		return helpers.ErrorHandler(ctx, "Cannot get tasks", fiber.StatusBadRequest)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"task": task,
	})
}

func DeleteTask(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	deletedCount, err := models.DeleteOne(id)
	if err != nil {
		return helpers.ErrorHandler(ctx, "Error deleting task", fiber.StatusConflict)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Deleted task successfully",
		"COUNT":   deletedCount,
	})
}

func UpdateTask(ctx *fiber.Ctx) error {
	var task models.Tasks
	user := ctx.Locals("user").(models.User)
	id := ctx.Params("id")

	if err := ctx.BodyParser(&task); err != nil {
		return helpers.ErrorHandler(ctx, "Error duruing task parsing", fiber.StatusBadRequest)
	}
	if task.Title == "" || task.Description == "" {
		return helpers.ErrorHandler(ctx, "Provide all fields", fiber.StatusBadRequest)
	}
	task.UpdatedAt = time.Now()
	task.User = user

	updateCount, err := models.UpdateOne(id, task)
	if err != nil {
		return helpers.ErrorHandler(ctx, "Error During UPdating task", fiber.StatusBadRequest)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Upadated task successfully",
		"Updated": updateCount,
	})
}
