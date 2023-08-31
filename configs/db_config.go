package configs

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func SetupDatabase() (*sql.DB, error) {
	//dbUsername := os.Getenv("DB_USERNAME")
	//dbPassword := os.Getenv("DB_PASSWORD")
	//dbName := os.Getenv("DB_NAME")

	//dbURL := fmt.Sprintf("user=%s password=%s dbname=%s", dbUsername, dbPassword, dbName)
	dbURL := fmt.Sprint("user=postgres password=root123 dbname=go_test host=localhost port=5432 sslmode=disable")
	fmt.Println(dbURL)
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}

	fmt.Println("Connected to the database")
	return db, nil
}
