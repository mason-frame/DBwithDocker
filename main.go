package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

//set up connection to db
const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "Chief0824"
	dbname   = "masonpractice"
)

func main() {
	//link to db in postgres
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// define columns in DB
	var name string
	var id string
	var age int

	// Query to get all results in the table
	rows, err := db.Query("SELECT id, name, age FROM information")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate through all the results
	for rows.Next() {
		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id %s\nmy name is %s\nand age is %d\n", id, name, age)
	}

	// Query to return at most one row
	row := db.QueryRow("SELECT id, name, age FROM information WHERE age=23")
	switch err := row.Scan(&id, &name, &age); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id, name, age)
	default:
		log.Fatal(err)
	}
	fmt.Println()

}
