package controller

import (
	"github.com/fernandogfraga/Golang_projetos/fiber_crud/src/service"
	"github.com/gofiber/fiber/v2"
)

func CreateProducts(c *fiber.Ctx) error {
	service.Example()
	return c.SendString("Hello World!")
}

func GetProducts(c *fiber.Ctx) error {
	service.Example()
	return c.SendString("Hello World!")
}

func UpdateProducts(c *fiber.Ctx) error {
	service.Example()
	return c.SendString("Hello World!")
}

func DeleteProducts(c *fiber.Ctx) error {
	service.Example()
	return c.SendString("Hello World!")
}
