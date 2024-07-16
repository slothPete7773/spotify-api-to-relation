package main

import (
	"log"
	"spotify-relation/repository"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// type Test struct {
// 	Name string `json:"name"`
// 	Age  uint   `json:"age"`
// }

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

	// "Open test file"
	// file, err := os.Open("data/test_duplicate_update.json")
	// if err != nil {
	// 	panic(err)
	// }

	// recentlyPlayedRecords := &model.RecentlyPlayedRecords{}

	// jsonParser := json.NewDecoder(file)
	// if err = jsonParser.Decode(&recentlyPlayedRecords); err != nil {
	// 	log.Fatal("parsing config file", err.Error())
	// }

	// // fmt.Printf("%v \n", recentlyPlayedRecords.Items[0].Track)

	// for _, activity := range recentlyPlayedRecords.Items {
	// 	// fmt.Printf("%v\n\n", activity)

	// 	artists := []db.Artist{}
	// 	for _, a := range activity.Track.Artists {
	// 		_artist := db.Artist{
	// 			Href:        a.Href,
	// 			ID:          a.ID,
	// 			Name:        a.Name,
	// 			ExternalUrl: a.ExternalUrls.Spotify,
	// 		}
	// 		artists = append(artists, _artist)
	// 		fmt.Printf("%v\n", _artist)

	// 	}

	// }
}
