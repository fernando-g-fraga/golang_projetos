package controller

import (
	"fmt"
	"log"
	"strconv"

	"github.com/fernandogfraga/Golang_projetos/fiber_crud/src/model"
	"github.com/fernandogfraga/Golang_projetos/fiber_crud/src/service"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func Controller_Update_Product(c *fiber.Ctx, db *pgx.Conn) error {
	// Initialize a new Product model to hold the parsed data
	product := new(model.Product)
	err := c.BodyParser(product)
	if err != nil {
		// Log the error and return a 400 Bad Request response
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid product data")
	}

	// Convert the id parameter from the URL to an integer
	convertedID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		// Log the error and return a 400 Bad Request response
		log.Printf("Error converting product id: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid product ID")
	}

	// Assign the converted ID to the product struct
	product.ID = convertedID

	// Log the parsed product details (optional, for debugging)
	fmt.Println(product)

	// Call the service function to update the product in the database
	err = service.Service_Update_Product(*product, db)
	if err != nil {
		// Log the error and return a 500 Internal Server Error response
		log.Printf("Error updating product with ID %d: %v", product.ID, err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error updating the product")
	}

	// Return a success message if everything went well
	return c.Status(fiber.StatusOK).SendString("Product updated successfully")
}

func Get_Products(c *fiber.Ctx, db *pgx.Conn) error {
	product := new(model.Product)

	products, err := service.Service_Get_Products(*product, db)
	if err != nil {
		c.Status(500).SendString("Error getting the products.")
	}
	return c.Status(200).JSON(products)
}

func Create_Product(c *fiber.Ctx, db *pgx.Conn) error {
	product := new(model.Product)
	err := c.BodyParser(product)
	if err != nil {
		log.Fatal(err)
	}

	id := service.Service_Post_Product(*product, db)

	return c.Status(200).JSON(map[string]string{"Message": fmt.Sprintf("Product Created with ID: %v", id)})
}

func Delete_product(c *fiber.Ctx, db *pgx.Conn) error {
	id := c.Params("id")
	service.Delete_product(id, db)
	return c.Status(200).JSON("ID deleted")
}
