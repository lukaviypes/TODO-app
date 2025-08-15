package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type DataBase struct {
	Db *sql.DB
}

func ConnectDb() *DataBase {
	connStr := "user=username dbname=mydb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return &DataBase{Db: db}
}
