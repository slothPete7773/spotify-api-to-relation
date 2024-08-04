package repository

import (
	"fmt"
	"spotify-relation/source"
	"time"

	"github.com/jmoiron/sqlx"
)

type trackRepositoryPgDB struct {
	db *sqlx.DB
}

func NewTrackRepositoryPgDB(db *sqlx.DB) TrackRepository {
	return trackRepositoryPgDB{db: db}
}

func (t trackRepositoryPgDB) IsSameWithExisting(sTrack *source.Track, eTrack *Track) bool {

	return (sTrack.ExternalUrls.Spotify == eTrack.ExternalUrl) &&
		(sTrack.Explicit == eTrack.Explicit) &&
		(sTrack.IsLocal == eTrack.IsLocal) &&
		(sTrack.Popularity == eTrack.Popularity)
}

func (t trackRepositoryPgDB) Upsert(track *source.Track) error {
	// Check if the track already exists
	existingTrack, err := t.GetById(track.ID)
	_ = existingTrack

	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil
	}

	if err != nil && err.Error() == "sql: no rows in result set" {
		err = t.Create(track)
		if err != nil {
			return fmt.Errorf("error creating track: %v", err)
		}
	} else {
		if t.IsSameWithExisting(track, existingTrack) == false {
			err = t.Update(track)
			if err != nil {
				return fmt.Errorf("error updating track: %v", err)
			}
			fmt.Printf("Updated Track ID: %v\n", track.ID)
		}

	}

	return nil
}

func (t trackRepositoryPgDB) GetAll() ([]Track, error) {
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
func (t trackRepositoryPgDB) GetById(trackId string) (*Track, error) {
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
		WHERE id = $1
	`
	err := t.db.Get(&track, query, trackId)
	if err != nil {
		return nil, err
	}
	return &track, nil
}
func (t trackRepositoryPgDB) Create(track *source.Track) error {

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
		, created_at
		, updated_at
	) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $12)
	`
	_, err := t.db.Exec(
		query,
		track.ID,
		track.Name,
		track.DurationMs,
		track.DiscNumber,
		track.ExternalUrls.Spotify,
		track.Explicit,
		track.IsLocal,
		track.Popularity,
		track.PreviewURL,
		track.TrackNumber,
		track.Album.ID,
		time.Now(),
	)
	if err != nil {
		return err
	}

	for _, artist := range track.Artists {
		query := `
		INSERT INTO track_artists (
			track_id
			, artist_id
		) VALUES ($1, $2)
		`
		_, err := t.db.Exec(query, track.ID, artist.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
func (t trackRepositoryPgDB) Update(track *source.Track) error {
	query := `
	UPDATE tracks SET
		name = $1
		, duration_ms = $2
		, disc_number = $3
		, external_url = $4
		, explicit = $5
		, is_local = $6
		, popularity = $7
		, preview_url = $8
		, track_number = $9
		, updated_at = $10
	WHERE id = $11
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
		time.Now(),
		track.ID,
	)
	if err != nil {
		return err
	}
	return nil
}
func (t trackRepositoryPgDB) IsExists(trackId string) bool {
	track := Track{}
	query := `
		SELECT
			id
		FROM tracks
		WHERE id = $1
	`

	err := t.db.Get(&track, query, trackId)
	if err != nil {
		return false
	}
	return true
}
