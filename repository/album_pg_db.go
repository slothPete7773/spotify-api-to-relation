package repository

import (
	"fmt"
	"spotify-relation/source"
	"time"

	"github.com/jmoiron/sqlx"
)

type albumRepositoryPgDB struct {
	db *sqlx.DB
}

func NewAlbumRepositoryPgDB(db *sqlx.DB) AlbumRepository {
	return albumRepositoryPgDB{db: db}
}

func (a albumRepositoryPgDB) GetAll() ([]Album, error) {
	albums := []Album{}
	query := `
		SELECT 
			id
			, name
			, album_type
			, release_date
			, release_date_precision
			, total_tracks
			, external_url
		FROM albums
	`
	err := a.db.Select(&albums, query)
	if err != nil {
		return nil, err
	}

	return albums, nil
}
func (a albumRepositoryPgDB) GetById(albumId string) (*Album, error) {
	album := Album{}
	query := `
		SELECT 
			id
			, name
			, album_type
			, release_date
			, release_date_precision
			, total_tracks
			, external_url
		FROM albums
		WHERE id = $1 
	`
	err := a.db.Get(&album, query, albumId)
	if err != nil {
		return nil, err
	}
	return &album, nil
}
func (a albumRepositoryPgDB) Create(album *source.Album) error {
	query := `
	INSERT INTO albums (
		id
		, name
		, album_type
		, release_date
		, release_date_precision
		, total_tracks
		, external_url
		, created_at
		, updated_at
	) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $8)`
	_, err := a.db.Exec(
		query,
		album.ID,
		album.Name,
		album.AlbumType,
		album.ReleaseDate,
		album.ReleaseDatePrecision,
		album.TotalTracks,
		album.ExternalUrls.Spotify,
		time.Now(),
	)
	if err != nil {
		return err
	}

	for _, artist := range album.Artists {
		fmt.Printf("AlbumID: %v\nArtistID: %v\n", album.ID, artist.ID)
		query := `
	INSERT INTO album_artists (
		album_id
		, artist_id
		, created_at
		, updated_at
	) VALUES ( $1, $2, $3, $3)`
		_, err := a.db.Exec(
			query,
			album.ID,
			artist.ID,
			time.Now(),
		)
		if err != nil {
			return err
		}
	}

	for _, img := range album.Images {
		query := `
	INSERT INTO album_images (
		album_id
		, image_id
		, created_at
		, updated_at
	) VALUES ( $1, $2, $3, $3)`
		_, err := a.db.Exec(
			query,
			album.ID,
			img.URL,
			time.Now(),
		)
		if err != nil {
			return err
		}
	}

	if err != nil {
		return err
	}
	return nil
}
func (a albumRepositoryPgDB) Update(album *source.Album) error {
	query := `
	UPDATE albums SET
		name = $1
		, album_type = $2
		, release_date = $3
		, release_date_precision = $4
		, total_tracks = $5
		, external_url = $6
		, updated_at = $7
	WHERE id = $8`
	_, err := a.db.Exec(
		query,
		album.Name,
		album.AlbumType,
		album.ReleaseDate,
		album.ReleaseDatePrecision,
		album.TotalTracks,
		album.ExternalUrls.Spotify,
		time.Now(),
		album.ID,
	)
	if err != nil {
		return err
	}
	return nil
}
func (a albumRepositoryPgDB) IsExists(albumId string) bool {
	album := Album{}
	query := `
		SELECT 
			id
		FROM albums
		WHERE id = $1
		`
	err := a.db.Get(&album, query, albumId)
	if err != nil {
		return false
	}

	return true
}
