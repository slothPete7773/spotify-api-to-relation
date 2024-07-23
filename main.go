package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"spotify-relation/repository"
	"spotify-relation/source"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sqlx.Open("sqlite3", "./spotify_data.db?mode=rwc")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	migrate()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	artistRepository := repository.NewArtistRepositorySQLiteDB(db)
	_ = artistRepository

	imageRepository := repository.NewImageRepositorySQLiteDB(db)
	_ = imageRepository

	albumRepository := repository.NewAlbumRepositoryDB(db)
	_ = albumRepository

	trackRepository := repository.NewTrackRepositoryDB(db)
	_ = trackRepository

	activityRepository := repository.NewActivityRepositoryDB(db)
	_ = activityRepository

	// "Open test file"
	// file, err := os.Open("data/test_duplicate_update.json")
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

		if isTrackAlreadyExists := trackRepository.IsExists(activity.Track.ID); isTrackAlreadyExists == false {
			fmt.Printf("Track ID: %v is not exists, creating...\n", activity.Track.ID)
			err = trackRepository.Create(&activity.Track)
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
