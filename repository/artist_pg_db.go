package repository

import (
	"spotify-relation/source"
	"time"

	"github.com/jmoiron/sqlx"
)

type artistRepositorySQLitePgDB struct {
	db *sqlx.DB
}

func NewArtistRepositorySQLitePgDB(db *sqlx.DB) ArtistRepository {
	return artistRepositorySQLitePgDB{
		db: db,
	}
}

func (a artistRepositorySQLitePgDB) IsExists(artistId string) bool {
	artist := Artist{}
	query := `
		SELECT 
			id
		FROM artists
		WHERE id = $1
		`
	err := a.db.Get(&artist, query, artistId)
	if err != nil {
		return false
	}

	return true
}

func (a artistRepositorySQLitePgDB) GetAll() ([]Artist, error) {
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
func (a artistRepositorySQLitePgDB) GetById(artistId string) (*Artist, error) {
	artist := Artist{}
	query := `
		SELECT 
			id
			, external_url
			, name
			, checksum
		FROM artists WHERE id = $1`
	err := a.db.Get(&artist, query, artistId)
	if err != nil {
		return nil, err
	}
	return &artist, nil
}
func (a artistRepositorySQLitePgDB) Create(artist *source.Artist) error {
	query := `
	INSERT INTO artists (
		id
		, name
		, external_url
		, created_at
		, updated_at
	) VALUES ( $1, $2, $3 , $4, $4)`
	_, err := a.db.Exec(query, artist.ID, artist.Name, artist.ExternalUrls.Spotify, time.Now())
	if err != nil {
		return err
	}
	return nil
}
func (a artistRepositorySQLitePgDB) Update(artist *source.Artist) error {
	query := `
	UPDATE artists SET
		external_url = $1
		, name = $2
		, updated_at = $3
	WHERE id = $4`
	_, err := a.db.Exec(query, artist.ExternalUrls.Spotify, artist.Name, time.Now(), artist.ID)
	if err != nil {
		return err
	}
	return nil
}
