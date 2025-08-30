package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DataBase struct {
	Db *sql.DB
}

func ConnectDb(myuser, mypassword, mydatabase string) (*DataBase, error) {
	//connStr := "user=myuser password=mypassword dbname=mydatabase sslmode=disable"
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", myuser, mypassword, mydatabase)

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
