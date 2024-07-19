package repository

import (
	"spotify-relation/source"
	"time"

	"github.com/jmoiron/sqlx"
)

type activityRepositoryDB struct {
	db *sqlx.DB
}

func NewActivityRepositoryDB(db *sqlx.DB) ActivityRepository {
	return activityRepositoryDB{db: db}
}

func (a activityRepositoryDB) GetByTrack(trackId string) ([]Activity, error) {
	activities := []Activity{}
	query := `
		SELECT 
			track_id
			, played_at
			, type
			, external_url
		FROM activities
		WHERE track_id = ?
	`
	err := a.db.Select(&activities, query, trackId)
	if err != nil {
		return nil, err
	}
	return activities, nil
}
func (a activityRepositoryDB) Create(activity *source.Activity) error {
	query := `
		INSERT INTO activities (
		track_id
		, played_at
		, context_type
		, external_url
		) VALUES (?, ?, ?, ?)
	`

	_, err := a.db.Exec(
		query,
		activity.Track.ID,
		activity.PlayedAt,
		activity.Context.Type,
		activity.Context.ExternalUrls.Spotify,
	)
	if err != nil {
		return err
	}

	return nil
}

func (a activityRepositoryDB) IsExists(playedTimestamp time.Time) bool {
	activity := Activity{}
	query := `
		SELECT 
			played_at
		FROM activities
		WHERE played_at = ?
	`

	err := a.db.Get(&activity, query, playedTimestamp)
	if err != nil {
		return false
	}
	return true
}
