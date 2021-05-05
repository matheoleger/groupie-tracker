package handlers

import (
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

//Pour la page artist par défaut
type PageDataArtists struct {
	Artists []Artist
}

//Pour un artist sur lequel on a appuyé sur "Voir plus..."
type MoreInformationArtist struct {
	Artist   Artist
	Relation Relation
}

// Pour la recherche (correspond à un artist mais pas forcément à une recherche possible)
type SearchResult struct {
	Id           int
	Image        string
	Name         string
	FirstAlbum   string
	CreationDate int
	Members      []string
	Locations    []string
}

func GetAPI(pathAPI string) []byte {
	restAPI := "https://groupietrackers.herokuapp.com/api"

	response, err := http.Get(restAPI + pathAPI)

	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}
