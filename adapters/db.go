package adapters

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	if db == nil {
		connectDb()
	}
}

func GetDb() *sql.DB {
	return db
}

func connectDb() {
	var err error
	connectionString := createConnectionString()
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Successful connection to the database")
}

func createConnectionString() string {
	host, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {
		host = "localhost"
	}

	port, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		port = "5433"
	}

	user, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		user = "postgres"
	}

	password, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		password = "admin"
	}

	dbName, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {
		dbName = "app-users"
	}

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",user, password, dbName, host, port)
	log.Printf("Connection string: %s\n", connectionString)
	
	return connectionString
}
