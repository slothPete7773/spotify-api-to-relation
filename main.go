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
	// "Open test file"
	file, err := os.Open("data/test_duplicate_update.json")
	if err != nil {
		panic(err)
	}

	recentlyPlayedRecords := &source.RecentlyPlayedRecords{}

	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&recentlyPlayedRecords); err != nil {
		log.Fatal("parsing config file", err.Error())
	}

	// fmt.Printf("%v \n", recentlyPlayedRecords.Items[0].Track)

	for _, activity := range recentlyPlayedRecords.Items {
		// fmt.Printf("%v\n\n", activity)

		// artists := []db.Artist{}
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

		// fmt.Printf("Album ID: %v", activity.Track.Album.ID)
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

		break

		// _artist := repository.Artist{
		// 	ID:          a.ID,
		// 	Name:        a.Name,
		// 	ExternalUrl: a.ExternalUrls.Spotify,
		// }
		// artists = append(artists, _artist)
		// fmt.Printf("%v\n", _artist)

	}
}

// "Image: Add"
// err = imageRepository.Add(&repository.Image{
// 	Height: 1,
// 	Width:  1,
// 	Url:    "url.com",
// })
// if err != nil {
// 	log.Fatal(err)
// }

// "Create"
// a, err := artistRepository.GetById("hhh")
// err = artistRepository.Create(&repository.Artist{
// 	ID:          "hello-id",
// 	Name:        "Veerakit",
// 	ExternalUrl: "url-veera.co",
// })
// if err != nil {
// 	log.Fatal(err)
// }

// "GetAll"
// artists, err := artistRepository.GetAll()
// if err != nil {
// 	log.Fatal(err)
// }
// for _, a := range artists {
// 	fmt.Printf("%v \n", a)
// }

// "GetById"
// a, err := artistRepository.GetById("hello-id")
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Printf("%v \n", a)

// "Update"
// a.Name = "slothpete"
// err = artistRepository.Update(a)
// if err != nil {
// 	log.Fatal(err)
// }
