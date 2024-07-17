package repository

import (
	"errors"
	"fmt"
	"spotify-relation/source"

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

func (a artistRepositorySQLiteDB) IsExists(artistId string) bool {
	artist := Artist{}
	query := `
		SELECT 
			id
		FROM artists
		WHERE id = ?
		`
	err := a.db.Get(&artist, query, artistId)
	if err != nil {
		return false
	}

	// return artists, nil
	return true
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
func (a artistRepositorySQLiteDB) Create(artist *source.Artist) error {
	_, err := a.GetById(artist.ID)
	if err == nil {
		// return nil
		return errors.New("Data already exists.")
	}
	query := `
	INSERT INTO artists (
		id
		, name
		, external_url
	) VALUES ( ?, ?, ? )`
	_, err = a.db.Exec(query, artist.ID, artist.Name, artist.ExternalUrls.Spotify)
	if err != nil {
		return err
	}
	return nil
}
func (a artistRepositorySQLiteDB) Update(artist *source.Artist) error {
	query := `
	UPDATE artists SET
		external_url = ?
		, name = ?
	WHERE id = ?`
	// , updated_at = :updated_at
	// artist.UpdatedAt = time.Now()
	fmt.Printf("to update:\n%v", artist)
	_, err := a.db.Exec(query, artist.ExternalUrls.Spotify, artist.Name, artist.ID)
	if err != nil {
		return err
	}
	return nil
}
