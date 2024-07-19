package repository

import (
	"spotify-relation/source"

	"github.com/jmoiron/sqlx"
)

type trackRepositoryDB struct {
	db *sqlx.DB
}

func NewTrackRepositoryDB(db *sqlx.DB) TrackRepository {
	return trackRepositoryDB{db: db}
}

func (t trackRepositoryDB) GetAll() ([]Track, error) {
	tracks := []Track{}
	query := `
		SELECT
			id
			, name
			, duration_ms
			, disc_number
			, external_url
			, explicit
			, is_local
			, popularity
			, preview_url
			, track_number
			, album_id
		FROM tracks
	`
	err := t.db.Select(&tracks, query)
	if err != nil {
		return nil, err
	}
	return tracks, nil
}
func (t trackRepositoryDB) GetById(trackId string) (*Track, error) {
	track := Track{}
	query := `
		SELECT
			id
			, name
			, duration_ms
			, disc_number
			, external_url
			, explicit
			, is_local
			, popularity
			, preview_url
			, track_number
			, album_id
		FROM tracks
		WHERE id = ?
	`
	err := t.db.Get(&track, query, trackId)
	if err != nil {
		return nil, err
	}
	return &track, nil
}
func (t trackRepositoryDB) Create(track *source.Track) error {

	query := `
	INSERT INTO tracks (
		id
		, name
		, duration_ms
		, disc_number
		, external_url
		, explicit
		, is_local
		, popularity
		, preview_url
		, track_number
		, album_id
	) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := t.db.Exec(
		query,
		track.ID,
		track.Name,
		track.DurationMs,
		track.DiscNumber,
		track.ExternalUrls.Spotify, // NOTENOTENOTNE
		track.Explicit,
		track.IsLocal,
		track.Popularity,
		track.PreviewURL,
		track.TrackNumber,
		track.Album.ID,
	)
	if err != nil {
		return err
	}

	for _, artist := range track.Artists {
		query := `
		INSERT INTO track_artists (
			track_id
			, artist_id
		) VALUES (?, ?)
		`
		_, err := t.db.Exec(query, track.ID, artist.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
func (t trackRepositoryDB) Update(track *source.Track) error {
	query := `
	UPDATE tracks SET
		name = ?
		, duration_ms = ?
		, disc_number = ?
		, external_url = ?
		, explicit = ?
		, is_local = ?
		, popularity = ?
		, preview_url = ?
		, track_number = ?
	WHERE id = ?
		`
	// album_id
	_, err := t.db.Exec(
		query,
		track.Name,
		track.DurationMs,
		track.DiscNumber,
		track.ExternalUrls.Spotify,
		track.Explicit,
		track.IsLocal,
		track.Popularity,
		track.PreviewURL,
		track.TrackNumber,
		track.ID,
	)
	if err != nil {
		return err
	}
	return nil
}
func (t trackRepositoryDB) IsExists(trackId string) bool {
	track := Track{}
	query := `
		SELECT
			id
		FROM tracks
		WHERE id = ?
	`

	err := t.db.Get(&track, query, trackId)
	if err != nil {
		return false
	}
	return true
}
