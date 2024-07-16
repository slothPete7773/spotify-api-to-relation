package repository

import "time"

type Activity struct {
	TrackId     string    `db:"track_id"`
	Track       Track     `gorm:"foreignKey:TrackId"`
	PlayedAt    time.Time `db:"played_at"`
	ContextType string    `db:"type"`
	ExternalUrl string    `db:"external_url"`
}

type ActivityRepository interface {
	GetAll()
	GetByTrack()
	Create()
}
