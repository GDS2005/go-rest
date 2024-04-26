package postgres

import (
	"database/sql"
	"example/configs"
	"fmt"

	_ "github.com/lib/pq"
)

// Initialize database connection
func InitDB() (*sql.DB, error) {
	// Construct database connection string
	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configs.DBHost, configs.DBPort, configs.DBUser, configs.DBPassword, configs.DBName)

	// Open connection to PostgreSQL
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}

	// Check if connection is established
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
