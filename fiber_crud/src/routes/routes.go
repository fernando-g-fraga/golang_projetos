package routes

import (
	"github.com/fernandogfraga/Golang_projetos/fiber_crud/src/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func Setup_routes(app *fiber.App, db *pgx.Conn) {

	api := app.Group("/api")

	api.Put("/update/:id", func(c *fiber.Ctx) error {
		return controller.Controller_Update_Product(c, db)
	})

	api.Get("/product", func(c *fiber.Ctx) error {
		return controller.Get_Products(c, db)
	})
	api.Post("/product/create", func(c *fiber.Ctx) error {
		return controller.Create_Product(c, db)
	})
	api.Delete("/product/delete/:id", func(c *fiber.Ctx) error {
		return controller.Delete_product(c, db)
	})

}
