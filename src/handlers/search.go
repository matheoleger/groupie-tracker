package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
)

func Search(w http.ResponseWriter, r *http.Request) {
	searchValue := r.URL.Query().Get("search")

	fmt.Println(searchValue)

	artistsData := GetAPI("/artists")
	var artistsList []Artist
	json.Unmarshal(artistsData, &artistsList)

	locationsData := GetAPI("/locations")
	// var locationsList []Location
	var locationsList Location
	json.Unmarshal(locationsData, &locationsList)
	// fmt.Println(locationsList)

	var finalArtistsListData []SearchResult

	re := regexp.MustCompile(`(?i)` + searchValue)

	for _, v := range artistsList {

		sortingData := SearchResult{}

		if re.Match([]byte(v.Name)) {
			// finalArtistsData.NameOfArtistResult.NameOfArtist = append(finalArtistsData.NameOfArtistResult.NameOfArtist, v.Name)
			// finalArtistsData.NameOfArtistResult.Id = append(finalArtistsData.NameOfArtistResult.Id, v.ID)
			// finalArtistsData.NameOfArtistResult.Image = append(finalArtistsData.NameOfArtistResult.Image, v.Image)
			sortingData.Name = v.Name

		}

		if re.Match([]byte(v.FirstAlbum)) {
			// finalArtistsData.FirstAlbumResult.FirstAlbum = append(finalArtistsData.FirstAlbumResult.FirstAlbum, v.FirstAlbum)
			// finalArtistsData.FirstAlbumResult.Id = append(finalArtistsData.FirstAlbumResult.Id, v.ID)
			// finalArtistsData.FirstAlbumResult.Image = append(finalArtistsData.FirstAlbumResult.Image, v.Image)

			sortingData.FirstAlbum = v.FirstAlbum
		}

		if re.Match([]byte(strconv.Itoa(v.CreationDate))) {
			// finalArtistsData.CreationDateResult.CreationDate = append(finalArtistsData.CreationDateResult.CreationDate, v.CreationDate)
			// finalArtistsData.CreationDateResult.Id = append(finalArtistsData.CreationDateResult.Id, v.ID)
			// finalArtistsData.CreationDateResult.Image = append(finalArtistsData.CreationDateResult.Image, v.Image)

			sortingData.CreationDate = v.CreationDate
		}

		for _, member := range v.Members {
			if re.Match([]byte(member)) {
				sortingData.Members = append(sortingData.Members, member)
			}
		}

		if !reflect.DeepEqual(sortingData, SearchResult{}) {
			// fmt.Println("vide")
			sortingData.Image = v.Image
			sortingData.Id = v.ID
			finalArtistsListData = append(finalArtistsListData, sortingData)
		}
	}

	// fmt.Println(locationsList)

	for i, v := range locationsList.Index {

		sortingData := SearchResult{}
		// fmt.Println(("into the other loop"))
		for _, eachLoc := range v.Locations {
			// fmt.Println(("everything is not fine"))
			if re.Match([]byte(eachLoc)) {
				// fmt.Println(("everything is fine"))
				sortingData.Locations = append(sortingData.Locations, eachLoc)
			}
		}

		if !reflect.DeepEqual(sortingData, SearchResult{}) {

			sortingData.Image = artistsList[i].Image
			sortingData.Id = artistsList[i].ID
			finalArtistsListData = append(finalArtistsListData, sortingData)
		}
	}

	finalArtistsResult, err := json.Marshal(&finalArtistsListData)

	if err != nil {
		log.Println(err)
	}

	w.Write(finalArtistsResult)
}
