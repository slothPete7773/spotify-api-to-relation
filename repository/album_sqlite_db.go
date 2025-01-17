package repository

import (
	"spotify-relation/source"

	"github.com/jmoiron/sqlx"
)

type albumRepositoryDB struct {
	db *sqlx.DB
}

func NewAlbumRepositoryDB(db *sqlx.DB) AlbumRepository {
	return albumRepositoryDB{db: db}
}

func (a albumRepositoryDB) GetAll() ([]Album, error) {
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
func (a albumRepositoryDB) GetById(albumId string) (*Album, error) {
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
		WHERE id = ? 
	`
	err := a.db.Get(&album, query, albumId)
	if err != nil {
		return nil, err
	}
	return &album, nil
}
func (a albumRepositoryDB) Create(album *source.Album) error {
	query := `
	INSERT INTO albums (
		id
		, name
		, album_type
		, release_date
		, release_date_precision
		, total_tracks
		, external_url
	) VALUES ( ?, ?, ?, ?, ?, ?, ? )`
	_, err := a.db.Exec(
		query,
		album.ID,
		album.Name,
		album.AlbumType,
		album.ReleaseDate,
		album.ReleaseDatePrecision,
		album.TotalTracks,
		album.ExternalUrls.Spotify,
	)

	for _, img := range album.Images {
		query := `
	INSERT INTO album_images (
		album_id
		, image_id
	) VALUES ( ?, ?)`
		_, err := a.db.Exec(
			query,
			album.ID,
			img.URL,
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
func (a albumRepositoryDB) Update(album *source.Album) error {
	query := `
	UPDATE albums SET
		name = ?
		, album_type = ?
		, release_date = ?
		, release_date_precision = ?
		, total_tracks = ?
		, external_url = ?
	WHERE id = ?`
	_, err := a.db.Exec(
		query,
		album.Name,
		album.AlbumType,
		album.ReleaseDate,
		album.ReleaseDatePrecision,
		album.TotalTracks,
		album.ExternalUrls.Spotify,
		album.ID,
	)
	if err != nil {
		return err
	}
	return nil
}
func (a albumRepositoryDB) IsExists(albumId string) bool {
	album := Album{}
	query := `
		SELECT 
			id
		FROM albums
		WHERE id = ?
		`
	err := a.db.Get(&album, query, albumId)
	if err != nil {
		return false
	}

	return true
}
