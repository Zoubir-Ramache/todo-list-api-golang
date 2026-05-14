package main

import (
	"ramache-zoubir/todo-list-api-golang/db"
	"ramache-zoubir/todo-list-api-golang/routes"
	"strconv"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
)


const PORT=8000
func main(){
	app:=fiber.New()
	routes.SetupMiddleware(app)	
	routes.SetupRoutes(app)
	db.Connect()
	app.Listen(":"+strconv.Itoa(PORT))
}