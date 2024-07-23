package repository

import (
	"spotify-relation/source"

	"github.com/jmoiron/sqlx"
)

type imageRepositorySQLiteDB struct {
	db *sqlx.DB
}

func NewImageRepositorySQLiteDB(db *sqlx.DB) ImageRepository {
	return imageRepositorySQLiteDB{db: db}
}

func (i imageRepositorySQLiteDB) IsExists(imageUrl string) bool {
	img := Image{}
	query := `
		SELECT 
			url
		FROM images
		WHERE url = ?
		`
	err := i.db.Get(&img, query, imageUrl)
	if err != nil {
		return false
	}

	// return artists, nil
	return true
}

func (i imageRepositorySQLiteDB) Add(img *source.Image) error {
	query := `
	INSERT INTO images (
		height
		, width
		, url
	) VALUES (
		?
		, ?
		, ?
	)`

	_, err := i.db.Exec(query, img.Height, img.Width, img.URL)
	if err != nil {
		return err
	}
	return nil
}
