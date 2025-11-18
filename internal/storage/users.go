package storage

import "errors"

//type User struct {
//	Id           int64
//	Name         string
//	HashPassword string
//}

func (db *DataBase) CreateUser(name, hashpassword string) error {
	res, err := db.Db.Exec("INSERT INTO users (name, hashpassword) VALUES ($1, $2) ", name, hashpassword)
	if err != nil {

		return errors.New("error inserting user")
	}
	id, err := res.RowsAffected()
	if err != nil {
		return errors.New("error inserting user")
	}
	if id <= 0 {
		return errors.New("user already exists")
	}
	return nil
}

func (db *DataBase) GetUser(name string) (string, error) {
	var HashPass string
	err := db.Db.QueryRow("SELECT hashpassword FROM users WHERE name = $1", name).Scan(&HashPass)
	if err != nil {
		return "", err
	}
	if HashPass == "" {
		return "", errors.New("user not found")
	}
	return HashPass, nil
}

//func (db *DataBase) DeleteUser(name, hashpassword string) error {}
