package handler

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func getDBConnection() (*sql.DB, error) {

	configuration := GetConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configuration.DB_HOST[0], configuration.DB_PORT[0], configuration.DB_USER[0], configuration.DB_PASS[0], configuration.DB_NAME[0])
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Error GetDBConnection() sql.Open:", err)
		os.Exit(3)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error GetDBConnection() sql.Ping:", err)
		os.Exit(3)
	}
	return db, nil
}

func CloseDB(db *sql.DB) error {
	return db.Close()
}
