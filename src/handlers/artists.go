package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Artists(w http.ResponseWriter, r *http.Request) {

	artistsData := GetAPI("/artists")
	var artistsList []Artist

	json.Unmarshal(artistsData, &artistsList)
	pageArtists := PageDataArtists{Artists: artistsList}

	if r.URL.Path == "/artists/" {

		//ParseFiles très important : on doit y mettre toutes les pages qui doivent être chargé
		//et donc choisir en fonction de celle qui doit l'être

		files := findPathFiles("./templates/artists.html")
		ts, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}

		ts.Execute(w, pageArtists)

	} else if r.URL.Path == "/artists/showmore" {

		id := r.URL.Query().Get("id") //le paramètre "id" dans l'URL
		if id != "" {

			//the artist global data
			idInt, err := strconv.Atoi(id)
			if err != nil {
				log.Println(err)
			}
			theArtist := artistsList[idInt-1]

			//the artist relation (dates et locations)
			relationData := GetAPI("/relation/" + id)
			var relationDatesLocations Relation
			json.Unmarshal(relationData, &relationDatesLocations)

			moreInformationStruct := MoreInformationArtist{Artist: theArtist, Relation: relationDatesLocations}

			moreInformationPage, err := json.Marshal(&moreInformationStruct)

			if err != nil {
				log.Println(err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(moreInformationPage)
		}

	} else if r.URL.Path == "/artists/filter" { //for give the filtered artist
		newFilteredArtistsData := Filter(artistsList, r)
		newFilteredArtists, err := json.Marshal(&newFilteredArtistsData)

		if err != nil {
			log.Println(err)
		}

		w.Write(newFilteredArtists)
	} else {
		http.Error(w, "Page not found", 404)
	}

}

func Filter(artistsList []Artist, r *http.Request) []Artist {

	var filterArtist []Artist

	artist := r.URL.Query().Get("artist")
	group := r.URL.Query().Get("group")
	members := r.URL.Query().Get("members")
	creationDates := r.URL.Query().Get("creationDates")
	firstAlbum := r.URL.Query().Get("firstAlbum")

	for _, v := range artistsList {
		var isRight bool = true

		//filtre checkbox
		if artist == "true" && group == "false" {
			isRight = FilterArtist(v)
		} else if artist == "false" && group == "true" {
			isRight = FilterGroup(v)
		} else if artist == "true" && group == "true" {
			isRight = true
		} else {
			continue
		}

		if !isRight {
			continue
		}

		//filtre members (range)
		if !FilterMembers(v, members) {
			continue
		}

		//filtre creationDates (select)
		if !FilterCreationDate(v, creationDates) {
			continue
		}

		//filtre firstAlbum (select)
		if !FilterFirstAlbum(v, firstAlbum) {
			continue
		}

		//ajouter au tableau final si la condition est respecté
		if isRight {
			filterArtist = append(filterArtist, v)
		}
	}

	return filterArtist
}

func FilterArtist(element Artist) bool {
	if len(element.Members) == 1 {
		return true
	} else {
		return false
	}
}

func FilterGroup(element Artist) bool {
	if len(element.Members) > 1 {
		return true
	} else {
		return false
	}
}

func FilterMembers(element Artist, members string) bool {
	nbrOfMem, err := strconv.Atoi(members)
	if err != nil {
		log.Println(err)
		return false
	}

	if len(element.Members) >= nbrOfMem || nbrOfMem == 0 {
		return true
	} else {
		return false
	}
}

func FilterCreationDate(element Artist, creationDate string) bool {

	nbrDate, err := strconv.Atoi(creationDate)
	if err != nil {
		log.Println(err)
		return false
	}

	if (element.CreationDate >= nbrDate && element.CreationDate <= (nbrDate+10)) || nbrDate == 0 {
		return true
	} else {
		return false
	}
}

func FilterFirstAlbum(element Artist, firstAlbum string) bool {
	nbrDatesAlbum, err := strconv.Atoi(firstAlbum)
	if err != nil {
		log.Println(err)
		return false
	}

	nbrDatesElement, err := strconv.Atoi(strings.Split(element.FirstAlbum, "-")[2])
	if err != nil {
		log.Println(err)
		return false
	}

	if (nbrDatesElement >= nbrDatesAlbum && nbrDatesElement <= (nbrDatesAlbum+10)) || nbrDatesAlbum == 0 {
		return true
	} else {
		return false
	}
}
