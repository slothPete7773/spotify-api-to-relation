package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type artistRepositorySQLiteDB struct {
	db *sqlx.DB
}

func NewArtistRepositorySQLiteDB(db *sqlx.DB) ArtistRepository {
	// TODO: To implement.
	return artistRepositorySQLiteDB{
		db: db,
	}
}

func (a artistRepositorySQLiteDB) GetAll() ([]Artist, error) {
	artists := []Artist{}
	err := a.db.Select(&artists, `
		SELECT 
			id
			, external_url
			, name
			, checksum
		FROM artists
		`)
	if err != nil {
		return nil, err
	}

	return artists, nil

}
func (a artistRepositorySQLiteDB) GetById(artistId string) (*Artist, error) {
	artist := Artist{}
	query := `
		SELECT 
			id
			, external_url
			, name
			, checksum
		FROM artists WHERE id = ?`
	err := a.db.Get(&artist, query, artistId)
	if err != nil {
		return nil, err
	}
	return &artist, nil
}
func (a artistRepositorySQLiteDB) Create(artist *Artist) error {
	query := `
	INSERT INTO artists (
		id
		, name
		, external_url
	) VALUES (
		:id
		, :name
		, :externalurl
	)`

	_, err := a.db.NamedExec(query, artist)
	if err != nil {
		return err
	}
	return nil
}
func (a artistRepositorySQLiteDB) Update(artist *Artist) error {
	query := `UPDATE artists SET
		external_url = :external_url
		, name = :name
		WHERE id = :id
		`
	// , updated_at = :updated_at
	// artist.UpdatedAt = time.Now()
	fmt.Printf("to update:\n%v", artist)
	_, err := a.db.NamedExec(query, artist)
	if err != nil {
		return err
	}
	return nil
}
