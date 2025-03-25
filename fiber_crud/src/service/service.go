package service

import (
	"context"
	"fmt"

	"github.com/fernandogfraga/Golang_projetos/fiber_crud/src/model"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
)

func Service_Update_Product(product model.Product, db *pgx.Conn) error {

	if _, err := db.Exec(context.Background(), "UPDATE product SET name=$1,description=$2 WHERE id = $3", product.Name, product.Description, product.ID); err != nil {
		log.Error(err)
		return fmt.Errorf("Error accessing database")
	}
	return nil
}

func Service_Get_Products(product model.Product, db *pgx.Conn) ([]model.Product, error) {
	rows, err := db.Query(context.Background(), "SELECT * FROM product")
	if err != nil {
		log.Error(err)
	}
	var products []model.Product

	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description); err != nil {
			return nil, err
		}
		products = append(products, product)

	}
	return products, nil
}

func Service_Post_Product(product model.Product, db *pgx.Conn) int {
	row := db.QueryRow(context.Background(), "INSERT INTO product (name, description) VALUES ($1,$2) RETURNING id", product.Name, product.Description)
	var id int
	row.Scan(&id)
	return id
}

func Delete_product(id string, db *pgx.Conn) {
	_, err := db.Exec(context.Background(), "DELETE FROM product WHERE id=$1 RETURNING id;", id)

	if err != nil {
		log.Fatal(err)
	}

}
