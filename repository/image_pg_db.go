package repository

import (
	"spotify-relation/source"
	"time"

	"github.com/jmoiron/sqlx"
)

type imageRepositorySQLitePgDB struct {
	db *sqlx.DB
}

func NewImageRepositorySQLitePgDB(db *sqlx.DB) ImageRepository {
	return imageRepositorySQLitePgDB{db: db}
}

func (i imageRepositorySQLitePgDB) IsExists(imageUrl string) bool {
	img := Image{}
	query := `
		SELECT 
			url
		FROM images
		WHERE url = $1
		`
	err := i.db.Get(&img, query, imageUrl)
	if err != nil {
		return false
	}

	// return artists, nil
	return true
}

func (i imageRepositorySQLitePgDB) Add(img *source.Image) error {
	query := `
	INSERT INTO images (
		height
		, width
		, url
		, created_at
		, updated_at
	) VALUES (
		$1
		, $2
		, $3
		, $4
		, $4
	)`

	_, err := i.db.Exec(query, img.Height, img.Width, img.URL, time.Now())
	if err != nil {
		return err
	}
	return nil
}
