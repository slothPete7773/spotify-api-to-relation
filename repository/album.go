package repository

import "time"

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
	GetAll()
	GetById()
	Create()
	Update()
}

type AlbumArtists struct {
	AlbumId   string     `db:"album_id" gorm:"primaryKey"`
	ArtistId  string     `db:"artist_id" gorm:"primaryKey"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type AlbumImages struct {
	AlbumId   string     `db:"album_id" gorm:"primaryKey"`
	ImageId   string     `db:"image_id" gorm:"primaryKey"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
