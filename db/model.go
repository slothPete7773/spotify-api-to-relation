package db

import "time"

type Image struct {
	Height int    `db:"height"`
	URL    string `db:"url"`
	Width  int    `db:"width"`
}

type Album struct {
	AlbumType            string   `db:"album_type"`
	Artists              []Artist `db:"artists"`
	ExternalUrl          string   `db:"external_url"`
	ID                   string   `db:"id"`
	Images               []Image  `db:"images"`
	Name                 string   `db:"name"`
	ReleaseDate          string   `db:"release_date"`
	ReleaseDatePrecision string   `db:"release_date_precision"`
	TotalTracks          int      `db:"total_tracks"`
}

type Artist struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	ExternalUrl string `db:"external_url"`
}

type Track struct {
	Album       Album    `db:"album"`
	Artists     []Artist `db:"artists"`
	DiscNumber  int      `db:"disc_number"`
	DurationMs  int      `db:"duration_ms"`
	ExternalUrl string   `db:"external_url"`
	Explicit    bool     `db:"explicit"`
	ID          string   `db:"id"`
	IsLocal     bool     `db:"is_local"`
	Name        string   `db:"name"`
	Popularity  int      `db:"popularity"`
	PreviewURL  string   `db:"preview_url"`
	TrackNumber int      `db:"track_number"`
}

type Activity struct {
	Track       Track     `db:"track"`
	PlayedAt    time.Time `db:"played_at"`
	ContextType string    `db:"type"`
	ExternalUrl string    `db:"external_url"`
}
