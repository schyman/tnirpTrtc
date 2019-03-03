package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func startDatabase(config DatabaseConfig) {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		config.User, config.Password, config.Name, config.Host, config.Port))

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connecte to postgres database: " + config.Name)
	testQueries()
}

func testQueries() {
	fmt.Println("# Querying")
	rows, err := db.Query("SELECT * FROM project")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var project_id int
		var project_company_id int
		var project_name string
		err = rows.Scan(&project_id, &project_company_id, &project_name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("uid | username | department ")
		fmt.Printf("%3v | %8v | %6v \n", project_id, project_company_id, project_name)
	}
}
