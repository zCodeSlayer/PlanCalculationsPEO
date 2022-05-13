package postgres

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go-postgres/logger"
	"os"
)

func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error.Println("Error loading .env file")
	}
	db, err := sql.Open("postgres", os.Getenv("DB_CONNECTION"))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		logger.Error.Println(err)
	}

	if os.Getenv("F_DEBUG_DB") == "true" {
		logger.Info.Println("db successfully connected")
	}
	// return the connection
	return db
}
