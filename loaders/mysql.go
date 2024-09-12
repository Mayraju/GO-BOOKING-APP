package loaders

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(0.0.0.0:3307)/test")

	if err != nil {
		fmt.Printf(" error occurred: %s", err)
	}
	err = db.Ping()
	fmt.Println("In ping method=======")
	if err != nil {
		fmt.Println("Error connecting to database %v", err)
	} else {
		fmt.Printf("Successfully connected to database")
	}
	return db, err
}

func GetDB() *sql.DB {

	return db
}
