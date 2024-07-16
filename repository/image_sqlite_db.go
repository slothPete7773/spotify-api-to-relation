package repository

import "github.com/jmoiron/sqlx"

type imageRepositorySQLiteDB struct {
	db *sqlx.DB
}

func NewImageRepositorySQLiteDB(db *sqlx.DB) ImageRepository {
	return imageRepositorySQLiteDB{db: db}
}

func (i imageRepositorySQLiteDB) Add(img *Image) error {
	query := `
	INSERT INTO images (
		height
		, width
		, url
	) VALUES (
		:height
		, :width
		, :url
	)`

	_, err := i.db.NamedExec(query, img)
	if err != nil {
		return err
	}
	return nil
}
