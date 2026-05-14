package todos

import (
	"log"
	"ramache-zoubir/todo-list-api-golang/db"
	"ramache-zoubir/todo-list-api-golang/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllTodos(c *fiber.Ctx) error {
	var allTodos []models.Todo
	title, content := c.Query("title"), c.Query("content")

	db.DbConnection.Where("title LIKE ? AND Content LIKE ?", "%"+title+"%", "%"+content+"%").Find(&allTodos)

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

func DeleteTodo(c *fiber.Ctx) error {
	IDs := c.Query("id")
	id, err := strconv.Atoi(IDs)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
	}
	db.DbConnection.Delete(&models.Todo{}, id)
	return c.JSON(fiber.Map{"deleted ": "todo " + IDs})
}

func UpdateTodo(c *fiber.Ctx) error {
	var newTodo models.Todo
	id := c.Query("id")

	err := c.BodyParser(&newTodo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	IDunit64, err := strconv.ParseUint(id, 10, 64)
	newTodo.ID = uint(IDunit64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is  required",
		})
	}
	db.DbConnection.Save(&newTodo)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"updated": newTodo,
	})
}
