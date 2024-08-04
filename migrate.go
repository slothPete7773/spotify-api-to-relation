package main

import (
	"fmt"
	"log"
	"os"
	"spotify-relation/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func migrate() {

	// db, err := gorm.Open(sqlite.Open("spotify_data.db"), &gorm.Config{})
	// dsn_dev := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", os.Getenv("PG_DEV_HOST"), os.Getenv("PG_DEV_USERNAME"), os.Getenv("PG_DEV_PASSWORD"), os.Getenv("PG_DEV_DATABASE"), os.Getenv("PG_DEV_PORT"))
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", os.Getenv("PG_HOST"), os.Getenv("PG_USERNAME"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DATABASE"), os.Getenv("PG_PORT"))

	// Reference: https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
	// Solved Issue: Stuck prepared statement cache by default. Need to disalbe it.
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

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
