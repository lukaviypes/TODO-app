package storage

type Storage interface {
	InsertTask(title string) (int64, error)
}

func (database *DataBase) InsertTask(title string) (int64, error) {
	var Id int64
	err := database.Db.QueryRow("INSERT INTO my_table (title) VALUES ($1) RETURNING id", title).Scan(&Id)
	if err != nil {
		return 0, err
	}

	return Id, nil
}
