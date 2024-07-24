package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"spotify-relation/repository"
	"spotify-relation/source"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// DB: Sqlite3
	// db, err := sqlx.Open("sqlite3", "./spotify_data.db?mode=rwc")

	// DB: Postgres
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("PG_USERNAME"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_DATABASE"))
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
	}
	migrate()

	// artistRepository := repository.NewArtistRepositorySQLiteDB(db)
	artistRepository := repository.NewArtistRepositorySQLitePgDB(db)
	_ = artistRepository

	// imageRepository := repository.NewImageRepositorySQLiteDB(db)
	imageRepository := repository.NewImageRepositorySQLitePgDB(db)
	_ = imageRepository

	// albumRepository := repository.NewAlbumRepositoryDB(db)
	albumRepository := repository.NewAlbumRepositoryPgDB(db)
	_ = albumRepository

	// trackRepository := repository.NewTrackRepositoryDB(db)
	trackRepository := repository.NewTrackRepositoryPgDB(db)
	_ = trackRepository

	// activityRepository := repository.NewActivityRepositoryDB(db)
	activityRepository := repository.NewActivityRepositoryPgDB(db)
	_ = activityRepository

	// "Open test file"
	// file, err := os.Open("data/test_duplicate_update.json")
	// file, err := os.Open("data/test_single.json")
	file, err := os.Open("data/1716023519_spotify_recent_50.json")
	if err != nil {
		panic(err)
	}

	recentlyPlayedRecords := &source.RecentlyPlayedRecords{}

	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&recentlyPlayedRecords); err != nil {
		log.Fatal("parsing config file", err.Error())
	}

	for _, activity := range recentlyPlayedRecords.Items {

		for _, artist := range activity.Track.Artists {
			if isArtistAlreadyExist := artistRepository.IsExists(artist.ID); isArtistAlreadyExist == false {
				fmt.Printf("Artist ID: %v is not exists, creating...\n", artist.ID)
				err = artistRepository.Create(&artist)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		for _, img := range activity.Track.Album.Images {
			if isImgUrlAlreadyExists := imageRepository.IsExists(img.URL); isImgUrlAlreadyExists == false {
				fmt.Printf("Image url: %v is not exists, creating...\n", img.URL)
				err = imageRepository.Add(&img)
				if err != nil {
					log.Fatal(err)
				}
			}

		}

		if isAlbumAlreadyExists := albumRepository.IsExists(activity.Track.Album.ID); isAlbumAlreadyExists == false {
			fmt.Printf("Album ID: %v is not exists, creating...\n", activity.Track.Album.ID)
			err = albumRepository.Create(&activity.Track.Album)
			if err != nil {
				log.Fatal(err)
			}

		}

		trackRepository.Upsert(&activity.Track)

		if isActivityExists := activityRepository.IsExists(activity.PlayedAt); isActivityExists == false {
			fmt.Printf("Activity at %v is not exists, inserting..\n", isActivityExists)
			err = activityRepository.Create(&activity)
			if err != nil {
				log.Fatal(err)
			}
		}

	}
}
