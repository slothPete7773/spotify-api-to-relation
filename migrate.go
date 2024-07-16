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

	if err := db.AutoMigrate(&repository.Artist{}); err != nil {
		panic(err)
	}

}
