package repository

import (
	"spotify-relation/source"
	"time"
)

type Activity struct {
	TrackId     string    `db:"track_id"`
	Track       Track     `gorm:"foreignKey:TrackId"`
	PlayedAt    time.Time `db:"played_at"`
	ContextType string    `db:"context_type"`
	ExternalUrl string    `db:"external_url"`
}

type ActivityRepository interface {
	GetByTrack(string) ([]Activity, error)
	Create(*source.Activity) error
	IsExists(playedTimestamp time.Time) bool
}
