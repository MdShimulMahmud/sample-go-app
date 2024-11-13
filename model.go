package main

import (
	"database/sql"
	"fmt"
)

type product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

func getProducts(db *sql.DB) ([]product, error) {
	query := "SELECT * FROM products"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	products := []product{}
	for rows.Next() {
		var p product
		err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.Price, &p.Description)

		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (p *product) getProduct(db *sql.DB) error {
	query := fmt.Sprintf("SELECT ID, Name, Quantity, Price, Description FROM products WHERE id=%v", p.ID)
	row := db.QueryRow(query)
	err := row.Scan(&p.ID, &p.Name, &p.Quantity, &p.Price, &p.Description)
	if err != nil {
		return err
	}
	return nil
}

func (p *product) createProduct(db *sql.DB) error {
	query := "INSERT INTO products(name, quantity, price, description) VALUES (?, ?, ?, ?)"
	res, err := db.Exec(query, p.Name, p.Quantity, p.Price, p.Description)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = int(id)
	return nil
}
