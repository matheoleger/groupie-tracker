package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json: image`
	Name         string   `json:"name"`
	Members      []string `json:members`
	CreationDate int      `json:creationDate`
	FirstAlbum   string   `json:firstAlbum`
}

type Location struct {
	// ID        int      `json:"id"`
	// Locations []string `json: locations`
	Index []struct {
		Id        int      `json: id`
		Locations []string `json: locations`
	} `json: index`
}

type Date struct {
	Dates []string `json: dates`
}

type Relation struct {
	DatesLocations interface{} `json: datesLocations`
}

// type PageDataArtists struct {
// 	Artists       []Artist
// 	LocationsList Location
// 	Dates         Date
// }

type PageDataArtists struct {
	Artists []Artist
}

// type MoreInformationArtist struct {
// 	Artist        Artist
// 	LocationsList Location
// 	DatesList     Date
// }

type MoreInformationArtist struct {
	Artist   Artist
	Relation Relation
}

// type SearchResult struct {
// 	NameOfArtistResult struct {
// 		Id           []int
// 		Image        []string
// 		NameOfArtist []string
// 	}
// 	FirstAlbumResult struct {
// 		Id         []int
// 		Image      []string
// 		FirstAlbum []string
// 	}
// 	CreationDateResult struct {
// 		Id           []int
// 		Image        []string
// 		CreationDate []int
// 	}
// 	Locations []Location
// }

type SearchResult struct {
	Id           int
	Image        string
	Name         string
	FirstAlbum   string
	CreationDate int
	Members      []string
	Locations    []string
}

// type MoreInformationArtist struct {
// 	// Artist   Artist
// 	// Relation Relation
// 	More []interface{}
// }

func GetApi(w http.ResponseWriter, r *http.Request) {

	response, err := http.Get("https://groupietrackers.herokuapp.com" + r.URL.Path)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(responseData)
}

func GetAPI(pathAPI string) []byte {
	restAPI := "https://groupietrackers.herokuapp.com/api"

	response, err := http.Get(restAPI + pathAPI)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	// fmt.Println(string(responseData) + "coucou cest ici")
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}
