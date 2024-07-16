package main

import (
	"log"
	"spotify-relation/repository"

	// "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func migrate() {

	db, err := gorm.Open(sqlite.Open("spotify_data.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(
		&repository.Activity{},
		&repository.Album{},
		&repository.AlbumImages{},
		&repository.AlbumArtists{},
		&repository.Artist{},
		&repository.Image{},
		&repository.Track{},
		&repository.TrackArtists{},
	); err != nil {
		panic(err)
	}

}
