package repository

import (
	"spotify-relation/source"
	"time"
)

type Track struct {
	DiscNumber  int        `db:"disc_number"`
	DurationMs  int        `db:"duration_ms"`
	ExternalUrl string     `db:"external_url"`
	Explicit    bool       `db:"explicit"`
	ID          string     `db:"id"`
	IsLocal     bool       `db:"is_local"`
	Name        string     `db:"name"`
	Popularity  int        `db:"popularity"`
	PreviewURL  string     `db:"preview_url"`
	TrackNumber int        `db:"track_number"`
	AlbumId     string     `db:"album_id"`
	Album       Album      `gorm:"foreignKey:AlbumId"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
	// Artists     []Artist `db:"artists"`
}

type TrackRepository interface {
	GetAll() ([]Track, error)
	GetById(string) (*Track, error)
	Create(*source.Track) error
	Update(*source.Track) error
	IsExists(string) bool
	Upsert(*source.Track) error
	IsSameWithExisting(*source.Track, *Track) bool
}

type TrackArtists struct {
	TrackId   string     `db:"track_id" gorm:"primaryKey"`
	Track     Track      `gorm:"foreignKey:TrackId"`
	ArtistId  string     `db:"artist_id" gorm:"primaryKey"`
	Artist    Artist     `gorm:"foreignKey:ArtistId"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
