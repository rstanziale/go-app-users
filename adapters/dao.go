package adapters

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var dao DAO

type DAO struct {
	db *sql.DB
}

func (dao *DAO) New(db *sql.DB) {
	dao.db = db
}

func (dao *DAO) GetDb() *sql.DB {
	return dao.db
}

func init() {
	dao = DAO{}
	dao.New(connectDb())
}

func GetDAO() DAO {
	return dao
}

func connectDb() *sql.DB {
	connectionString := createConnectionString()

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Successful connection to the database")

	return db
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
