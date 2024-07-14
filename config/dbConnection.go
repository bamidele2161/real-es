package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"log"
)

//code to connect to database
type Database struct {
	Db *sql.DB
}
func NewDb () (*Database, error){
	db := &Database{}
	err := db.Connect()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *Database) Connect() error {
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUsername := os.Getenv("DB_USERNAME")
	dbHost := os.Getenv("DB_HOST")
	 
	var err error
	connectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUsername, dbPassword, dbHost, dbName)
	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("checking error",err)
		log.Fatal(err)
		return err
	}
	d.Db = database
	return nil
}