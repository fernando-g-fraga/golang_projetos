package main

import (
	"github.com/fernandogfraga/Golang_projetos/fiber_crud/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.Setup_routes(app)

	app.Listen(":8080")

}
