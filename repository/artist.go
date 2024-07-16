package repository

import (
	"spotify-relation/source"
	"time"
)

type Artist struct {
	ID          string     `db:"id"`
	Name        string     `db:"name"`
	ExternalUrl string     `db:"external_url"`
	Checksum    *string    `db:"checksum"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}

type ArtistRepository interface {
	GetAll() ([]Artist, error)
	GetById(string) (*Artist, error)
	Create(*source.Artist) error
	Update(*source.Artist) error
}
