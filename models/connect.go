package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var con *sql.DB

// Connect initializes and returns a MySQL database connection
func Connect() (*sql.DB, error) {
	// Load environment variables
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Warning: Error loading .env file, using system environment variables")
	}

	// Get DB credentials from .env
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DATABASE")

	// Connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	// Test connection
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("database unreachable: %w", err)
	}

	fmt.Println("Connected to MySQL successfully!")
	con = db
	return db, nil
}
