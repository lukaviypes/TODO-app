package storage

type Storage interface {
	InsertTask(title string) (int64, error)
}

func (db *DataBase) InsertTask(title string) (int64, error) {
	var Id int64
	res, err := db.Db.Query("INSERT INTO tasks (title) VALUES ($1) RETURNING id", title)
	if err != nil {
		return 0, err
	}
	if err = res.Scan(Id); err != nil {
		return 0, err
	}

	return Id, nil
}
