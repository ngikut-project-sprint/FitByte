package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // Import the PostgreSQL driver.
)

var db *sql.DB

// Init initializes the database connection.
func Init() {
	connStr := os.Getenv("DATABASE_URL") // Use environment variables for configuration.
	if connStr == "" {
		log.Fatal("DATABASE_URL is not set!")
	}

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Test the connection.
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	fmt.Println("Database successfully connected!")
}

// GetDB returns the database instance.
func GetDB() *sql.DB {
	if db == nil {
		log.Fatal("Database is not initialized. Did you forget to call database.Init()?")
	}
	return db
}

// Close closes the database connection.
func Close() {
	if db != nil {
		db.Close()
	}
}
