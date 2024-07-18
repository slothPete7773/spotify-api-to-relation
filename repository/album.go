package repository

import (
	"spotify-relation/source"
	"time"
)

type Album struct {
	ID                   string     `db:"id"`
	Name                 string     `db:"name"`
	AlbumType            string     `db:"album_type"`
	ReleaseDate          string     `db:"release_date"`
	ReleaseDatePrecision string     `db:"release_date_precision"`
	TotalTracks          int        `db:"total_tracks"`
	ExternalUrl          string     `db:"external_url"`
	CreatedAt            *time.Time `db:"created_at"`
	UpdatedAt            *time.Time `db:"updated_at"`
	DeletedAt            *time.Time `db:"deleted_at"`
}

type AlbumRepository interface {
	GetAll() ([]Album, error)
	GetById(string) (*Album, error)
	Create(*source.Album) error
	Update(*source.Album) error
	IsExists(string) bool
}

type AlbumArtists struct {
	AlbumId   string     `db:"album_id" gorm:"primaryKey"`
	Album     Album      `gorm:"foreignKey:AlbumId"`
	ArtistId  string     `db:"artist_id" gorm:"primaryKey"`
	Artist    Artist     `gorm:"foreignKey:ArtistId"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type AlbumImages struct {
	AlbumId   string     `db:"album_id" gorm:"primaryKey"`
	Album     Album      `gorm:"foreignKey:AlbumId"`
	ImageId   string     `db:"image_id" gorm:"primaryKey"`
	Image     Image      `gorm:"foreignKey:ImageId"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
