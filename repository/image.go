package repository

import (
	"spotify-relation/source"
	"time"
)

type Image struct {
	Height    int        `db:"height"`
	Width     int        `db:"width"`
	Url       string     `db:"url" gorm:"primaryKey"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type ImageRepository interface {
	Add(*source.Image) error
	IsExists(string) bool
}
