package routes

import (
	"ramache-zoubir/todo-list-api-golang/internal/todos"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

)

func SetupRoutes(app *fiber.App) {
	app.Get("/", todos.GetAllTodos)
	app.Post("/", todos.CreateTodo)
	app.Delete("/", todos.DeleteTodo)
	app.Patch("/" , todos.UpdateTodo)
}

func SetupMiddleware(app * fiber.App){
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
        AllowMethods: "GET, POST, PUT, DELETE",
	}))
}