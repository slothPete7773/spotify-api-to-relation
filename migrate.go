package main

import (
	"fmt"
	"log"
	"os"
	"spotify-relation/repository"

	// "github.com/mattn/go-sqlite3"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func migrate() {

	// db, err := gorm.Open(sqlite.Open("spotify_data.db"), &gorm.Config{})
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Bangkok"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", os.Getenv("PG_HOST"), os.Getenv("PG_USERNAME"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DATABASE"), os.Getenv("PG_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(
		&repository.Activity{},
		&repository.Album{},
		&repository.AlbumImages{},
		&repository.Artist{},
		&repository.Image{},
		&repository.Track{},
		&repository.TrackArtists{},
	); err != nil {
		panic(err)
	}
	fmt.Printf("Successfully migrated.")

}
