package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"spotify-relation/repository"
	"spotify-relation/source"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "credential/slothpete7773-warehouse.json")
}

func main() {
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

	projectId := "slothpete7773-warehouse"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal("Failed to create client: %v", err)
	}
	defer client.Close()
	bucketName := "spotify_recently_played_slothpete7773"
	storageClient := repository.NewStorageGCP(client, projectId, bucketName, "spotify-recently-played/")

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

	jsonFiles, err := ListAllJson(os.Getenv("LANDING_DIRECTORY"))
	if err != nil {
		log.Fatal(err)
	}

	for _, filepath := range jsonFiles {

		file, err := os.Open(filepath)

		if err != nil {
			panic(err)
		}

		recentlyPlayedRecords := &source.RecentlyPlayedRecords{}

		jsonParser := json.NewDecoder(file)
		if err = jsonParser.Decode(&recentlyPlayedRecords); err != nil {
			log.Fatal("parsing config file", err.Error())
		}

		for i, activity := range recentlyPlayedRecords.Items {
			fmt.Printf("[%v] Activity at %v\n", i, activity.PlayedAt)

			for _, artist := range activity.Track.Artists {
				if isArtistAlreadyExist := artistRepository.IsExists(artist.ID); isArtistAlreadyExist == false {
					// fmt.Printf("Artist ID: %v is not exists, creating...\n", artist.ID)
					err = artistRepository.Create(&artist)
					if err != nil {
						log.Fatal(err)
					}
				}
			}

			for _, img := range activity.Track.Album.Images {
				if isImgUrlAlreadyExists := imageRepository.IsExists(img.URL); isImgUrlAlreadyExists == false {
					// fmt.Printf("Image url: %v is not exists, creating...\n", img.URL)
					err = imageRepository.Add(&img)
					if err != nil {
						log.Fatal(err)
					}
				}

			}

			if isAlbumAlreadyExists := albumRepository.IsExists(activity.Track.Album.ID); isAlbumAlreadyExists == false {
				// fmt.Printf("Album ID: %v is not exists, creating...\n", activity.Track.Album.ID)
				err = albumRepository.Create(&activity.Track.Album)
				if err != nil {
					log.Fatal(err)
				}

			}

			trackRepository.Upsert(&activity.Track)

			if isActivityExists := activityRepository.IsExists(activity.PlayedAt); isActivityExists == false {
				// fmt.Printf("[%v] Activity at %v is not exists, inserting..\n", i, activity.PlayedAt)

				err = activityRepository.Create(&activity)
				if err != nil {
					log.Fatal(err)
				}
			}

		}

		temp_filepath := strings.Split(filepath, "/")
		filename := temp_filepath[len(temp_filepath)-1]

		storageClient.UploadFile(filepath, filename)
		if err != nil {
			log.Fatal(err)
		}
		err = os.Remove(filepath)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func ListAllJson(target_directory string) ([]string, error) {
	filepaths := []string{}

	err := filepath.WalkDir(target_directory, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && filepath.Ext(path) == ".json" {
			filepaths = append(filepaths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
		// log.Fatal(err)
	}
	return filepaths, nil
}
