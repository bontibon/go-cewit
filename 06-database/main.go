package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// START OMIT
	db, err := sql.Open("sqlite3", "06-database/chinook.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var name string
	var sum float64

	err = db.QueryRow(`
			SELECT
				Customer.FirstName || " " || Customer.LastName,
				Invoice.Total
			FROM Invoice
			LEFT JOIN Customer
				ON Invoice.CustomerId = Customer.CustomerId
			ORDER BY Total DESC
			LIMIT 1
		`).
		Scan(&name, &sum)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s had the largest order of $%.02f\n", name, sum)
	// END OMIT
}
