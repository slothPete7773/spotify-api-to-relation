package source

import "time"

type Image struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type Album struct {
	AlbumType            string      `json:"album_type"`
	Artists              []Artist    `json:"artists"`
	AvailableMarkets     []string    `json:"available_markets"`
	ExternalUrls         ExternalUrl `json:"external_urls"`
	Href                 string      `json:"href"`
	ID                   string      `json:"id"`
	Images               []Image     `json:"images"`
	Name                 string      `json:"name"`
	ReleaseDate          string      `json:"release_date"`
	ReleaseDatePrecision string      `json:"release_date_precision"`
	TotalTracks          int         `json:"total_tracks"`
	Type                 string      `json:"type"`
	URI                  string      `json:"uri"`
}

type Artist struct {
	ExternalUrls ExternalUrl `json:"external_urls"`
	Href         string      `json:"href"`
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	URI          string      `json:"uri"`
}

type Track struct {
	Album            Album       `json:"album"`
	Artists          []Artist    `json:"artists"`
	AvailableMarkets []string    `json:"available_markets"`
	DiscNumber       int         `json:"disc_number"`
	DurationMs       int         `json:"duration_ms"`
	Explicit         bool        `json:"explicit"`
	ExternalIds      ExternalId  `json:"external_ids"`
	ExternalUrls     ExternalUrl `json:"external_urls"`
	Href             string      `json:"href"`
	ID               string      `json:"id"`
	IsLocal          bool        `json:"is_local"`
	Name             string      `json:"name"`
	Popularity       int         `json:"popularity"`
	PreviewURL       string      `json:"preview_url"`
	TrackNumber      int         `json:"track_number"`
	Type             string      `json:"type"`
	URI              string      `json:"uri"`
}

type ExternalId struct {
	Isrc string `json:"isrc"`
}
type ExternalUrl struct {
	Spotify string `json:"spotify"`
}

type Context struct {
	Type         string      `json:"type"`
	ExternalUrls ExternalUrl `json:"external_urls"`
	Href         string      `json:"href"`
	URI          string      `json:"uri"`
}

type Activity struct {
	Track    Track     `json:"track"`
	PlayedAt time.Time `json:"played_at"`
	Context  Context   `json:"context"`
}

type Cursors struct {
	After  string `json:"after"`
	Before string `json:"before"`
}
type RecentlyPlayedRecords struct {
	Items   []Activity `json:"items"`
	Next    string     `json:"next"`
	Cursors Cursors    `json:"cursors"`
	Limit   int        `json:"limit"`
	Href    string     `json:"href"`
}
