package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DataBase struct {
	Db *sql.DB
}

func ConnectDb() (*DataBase, error) {
	connStr := "user=myuser password=mypassword dbname=mydatabase sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DataBase{Db: db}, nil
}

func (db *DataBase) CloseDb() {
	db.Db.Close()
}
