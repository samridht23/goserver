package utils

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
)

// db instance
var db *sql.DB

func ConnectDb() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading enviroment variable", err)
	}
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("SQLPORT"))
	if err != nil {
		log.Fatal("Port not loaded from env")
	}
	user := os.Getenv("USER")
	password := os.Getenv("SQLPASS")
	dbname := os.Getenv("DBNAME")

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
