package repository

import (
	"spotify-relation/source"
	"time"

	"github.com/jmoiron/sqlx"
)

type activityRepositoryPgDB struct {
	db *sqlx.DB
}

func NewActivityRepositoryPgDB(db *sqlx.DB) ActivityRepository {
	return activityRepositoryPgDB{db: db}
}

func (a activityRepositoryPgDB) GetByTrack(trackId string) ([]Activity, error) {
	activities := []Activity{}
	query := `
		SELECT 
			track_id
			, played_at
			, type
			, external_url
		FROM activities
		WHERE track_id = $1
	`
	err := a.db.Select(&activities, query, trackId)
	if err != nil {
		return nil, err
	}
	return activities, nil
}
func (a activityRepositoryPgDB) Create(activity *source.Activity) error {
	query := `
		INSERT INTO activities (
		track_id
		, played_at
		, context_type
		, external_url
		) VALUES ($1, $2, $3, $4)
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

func (a activityRepositoryPgDB) IsExists(playedTimestamp time.Time) bool {
	activity := Activity{}
	query := `
		SELECT 
			played_at
		FROM activities
		WHERE played_at = $1
	`

	err := a.db.Get(&activity, query, playedTimestamp)
	if err != nil {
		return false
	}
	return true
}
