package todos

import (
	"log"
	"ramache-zoubir/todo-list-api-golang/db"
	"ramache-zoubir/todo-list-api-golang/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllTodos(c *fiber.Ctx) error {
	var allTodos []models.Todo
	db.DbConnection.Find(&allTodos)

	return c.JSON(allTodos)

}

func CreateTodo(c *fiber.Ctx) error {

	var newTodo models.Todo
	if err := c.BodyParser(&newTodo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error ",
		})
	}
	db.DbConnection.Create(&newTodo)

	log.Output(1, newTodo.Title)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "created"})
}
