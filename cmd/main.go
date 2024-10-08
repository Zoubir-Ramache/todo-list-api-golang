package main

import (
	
	"strconv"
	"ramache-zoubir/todo-list-api-golang/routes"
	"github.com/gofiber/fiber/v2"
	"ramache-zoubir/todo-list-api-golang/db"
)


const PORT=8000
func main(){
	app:=fiber.New()
	routes.SetupRoutes(app)
	db.Connect()
	app.Listen(":"+strconv.Itoa(PORT))
}