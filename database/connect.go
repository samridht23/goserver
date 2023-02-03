package database

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"fmt"
	"log"
)

// db instance
var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "paradox"
	dbname   = "fibertest"
)

func ConnectDb() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Env not loaded", err)
	}
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Database connected")
	return nil
}
