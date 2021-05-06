package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
)

func Search(w http.ResponseWriter, r *http.Request) {
	searchValue := r.URL.Query().Get("search")

	//On récupère l'API "artist" puis on l'a "Unmarshal" afin de la mettre dans un tableau de Struct (struct basé sur un "Artist")
	artistsData := GetAPI("/artists")
	var artistsList []Artist
	json.Unmarshal(artistsData, &artistsList)

	//Pareil
	locationsData := GetAPI("/locations")
	var locationsList Location
	json.Unmarshal(locationsData, &locationsList)

	var finalArtistsListData []SearchResult

	//Regex de la recherche avec le "(?i)" qui correspond à une mise en Majuscule
	re := regexp.MustCompile(`(?i)` + searchValue)

	for _, v := range artistsList {

		sortingData := SearchResult{}

		if re.Match([]byte(v.Name)) {
			sortingData.Name = v.Name

		}

		if re.Match([]byte(v.FirstAlbum)) {
			sortingData.FirstAlbum = v.FirstAlbum
		}

		if re.Match([]byte(strconv.Itoa(v.CreationDate))) {
			sortingData.CreationDate = v.CreationDate
		}

		for _, member := range v.Members {
			if re.Match([]byte(member)) {
				sortingData.Members = append(sortingData.Members, member)
			}
		}

		//Vérification de si sortingData n'est pas vide
		if !reflect.DeepEqual(sortingData, SearchResult{}) {
			sortingData.Image = v.Image
			sortingData.Id = v.Id
			finalArtistsListData = append(finalArtistsListData, sortingData)
		}
	}

	for i, v := range locationsList.Index {

		sortingData := SearchResult{}

		for _, eachLoc := range v.Locations {
			if re.Match([]byte(eachLoc)) {
				sortingData.Locations = append(sortingData.Locations, eachLoc)
			}
		}

		if !reflect.DeepEqual(sortingData, SearchResult{}) {

			sortingData.Image = artistsList[i].Image
			sortingData.Id = artistsList[i].Id
			finalArtistsListData = append(finalArtistsListData, sortingData)
		}
	}

	finalArtistsResult, err := json.Marshal(&finalArtistsListData)

	if err != nil {
		log.Println(err)
	}

	w.Write(finalArtistsResult)
}
