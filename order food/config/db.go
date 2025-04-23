package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DbURL constructs the MySQL connection string from environment variables.
func DbURL() (connectionString string) {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	protocol := os.Getenv("DB_PROTOCOL")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connectionString = fmt.Sprintf("%s:%s@%s(%s:%s)/%s", user, password, protocol, host, port, dbName)
	return
}

// InitDB initializes and returns a new database connection using named return values.
func InitDB() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", DbURL())
	if err != nil {
		fmt.Println("Error opening database connection:", err)
		return
	}

	if err = db.Ping(); err != nil {
		fmt.Println("Database connection failed:", err)
		return
	}

	fmt.Println("Connected to MySQL successfully!")
	return
}