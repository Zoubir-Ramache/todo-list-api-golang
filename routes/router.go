package routes

import (
	"ramache-zoubir/todo-list-api-golang/internal/todos"

	"github.com/gofiber/fiber/v2"
)


func SetupRoutes( app * fiber.App){
	app.Get("/" , todos.GetAllTodos)
	app.Post("/" , todos.CreateTodo)
}