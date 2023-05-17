package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=password dbname=labora-proyect-1 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Iniciamos una transacción
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Hacemos algunas operaciones de base de datos dentro de la transacción.
	_, err = tx.Exec("INSERT INTO items (customer_name, order_date, product, quantity, price) VALUES ($1, $2, $3, $4, $5)",
		"Jhon Kent", "2023-05-18", "Lenovo", 1, 2000)
	if err != nil {
		// Si algo salió mal, hacemos un rollback de la transacción
		tx.Rollback()
		log.Fatal(err)
	}

	_, err = tx.Exec("UPDATE items SET price = $1 WHERE customer_name = $2 AND product = $3",
		2500, "John Doe", "Apple Macbook Pro")
	if err != nil {
		// Si algo salió mal, hacemos un rollback de la transacción
		tx.Rollback()
		log.Fatal(err)
	}

	// Si todo salió bien, hacemos commit de la transacción
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transacción completada exitosamente")
}
