package adapters

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DAO struct {
	db *sql.DB
}

var dao DAO

func init() {
	dao = DAO{}
	dao.connectDb()
}

func GetDAO() DAO {
	return dao
}

// Return database
func (dao *DAO) GetDb() *sql.DB {
	return dao.db
}

// Connect database
func (dao *DAO) connectDb() {
	connectionString := dao.createConnectionString()

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Successful connection to the database")

	dao.db = db
}

// Create the connection string
func (dao *DAO) createConnectionString() string {
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
