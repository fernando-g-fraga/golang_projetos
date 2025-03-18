package routes

import (
	"github.com/fernandogfraga/Golang_projetos/fiber_crud/src/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup_routes(app *fiber.App) {

	api := app.Group("/api")

	api.Post("/create", controller.CreateProducts)
	api.Get("/", controller.GetProducts)
	api.Put("/edit", controller.UpdateProducts)
	api.Delete("/delete", controller.DeleteProducts)

}
