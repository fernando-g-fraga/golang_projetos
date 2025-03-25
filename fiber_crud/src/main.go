package main

import (
	"github.com/fernandogfraga/Golang_projetos/fiber_crud/src/database"
	"github.com/fernandogfraga/Golang_projetos/fiber_crud/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db, err := database.ConectDB()
	if err != nil {
		panic("Error seting up the Db")
	}
	routes.Setup_routes(app, db)
	app.Listen(":8080")

}
