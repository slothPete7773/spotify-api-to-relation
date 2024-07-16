package repository

import "time"

type Track struct {
	DiscNumber  int    `db:"disc_number"`
	DurationMs  int    `db:"duration_ms"`
	ExternalUrl string `db:"external_url"`
	Explicit    bool   `db:"explicit"`
	ID          string `db:"id"`
	IsLocal     bool   `db:"is_local"`
	Name        string `db:"name"`
	Popularity  int    `db:"popularity"`
	PreviewURL  string `db:"preview_url"`
	TrackNumber int    `db:"track_number"`
	// Album       Album    `db:"album"`
	// Artists     []Artist `db:"artists"`
}

type TrackRepository interface {
	GetAll()
	GetById()
	Create()
	Update()
}

type TrackArtists struct {
	TrackId   string     `db:"track_id" gorm:"primaryKey"`
	ArtistId  string     `db:"artist_id" gorm:"primaryKey"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
