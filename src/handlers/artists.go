package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Artists(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL.Path)
	// dataJson := usefulFiles.GetApi()

	//ParseFiles très important : on doit y mettre toutes les pages qui doivent être chargé et donc choisir en fonction de celle qui doit l'être
	// le choix doit être fait avant le ParseFiles
	// files := findPathFiles("./templates/artists.html")
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	http.Error(w, "Internal Server Error", 500)
	// 	return
	// }

	artistsData := GetAPI("/artists")
	var artistsList []Artist
	// var anArtist
	json.Unmarshal(artistsData, &artistsList)
	pageArtists := PageDataArtists{Artists: artistsList}

	if r.URL.Path == "/artists/" {

		files := findPathFiles("./templates/artists.html")
		ts, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}

		// id := r.URL.Query().Get("id")
		// if id != "" {
		// 	fmt.Println(id)

		// 	//the artist global data
		// 	idInt, err := strconv.Atoi(id)
		// 	if err != nil {
		// 		log.Println(err)
		// 	}
		// 	theArtist := artistsList[idInt-1]

		// 	//the artist relation (dates et locations)
		// 	relationData := GetAPI("/relation/" + id)
		// 	var relationDatesLocations Relation
		// 	json.Unmarshal(relationData, &relationDatesLocations)

		// 	var testArray [2]interface{}
		// 	testArray[0] = theArtist
		// 	testArray[1] = relationDatesLocations
		// 	// moreInformationStruct := MoreInformationArtist{More: theArtist, Relation: relationDatesLocations}

		// 	// var moreInformationPage []byte
		// 	// moreInformationPage, err := json.Marshal(&moreInformationStruct)
		// 	moreInformationPage, err := json.Marshal(&testArray)

		// 	if err != nil {
		// 		log.Println(err)
		// 	}

		// 	fmt.Println(string(moreInformationPage))

		// 	w.Header().Set("Content-Type", "application/json")
		// 	w.Write(moreInformationPage)
		// }

		ts.Execute(w, pageArtists)

	} else if r.URL.Path == "/artists/showmore" {

		id := r.URL.Query().Get("id")
		if id != "" {
			fmt.Println(id)

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

			// var testArray [2]interface{}
			// testArray[0] = theArtist
			// testArray[1] = relationDatesLocations
			moreInformationStruct := MoreInformationArtist{Artist: theArtist, Relation: relationDatesLocations}

			// var moreInformationPage []byte
			moreInformationPage, err := json.Marshal(&moreInformationStruct)
			// moreInformationPage, err := json.Marshal(&testArray)

			if err != nil {
				log.Println(err)
			}

			fmt.Println(string(moreInformationPage))

			w.Header().Set("Content-Type", "application/json")
			w.Write(moreInformationPage)
		}

	} else if r.URL.Path == "/artists/filter" { //for give the filtered artist
		newFilteredArtistsData := Filter(artistsList, r)
		// fmt.Println(newFilteredArtistsData)
		newFilteredArtists, err := json.Marshal(&newFilteredArtistsData)

		if err != nil {
			log.Println(err)
		}

		w.Write(newFilteredArtists)
	} else {
		http.Error(w, "Page not found", 404)
	}
	// ts, err := template.ParseFiles("./templates/layout.html", "./templates/location.html")

}

func Filter(artistsList []Artist, r *http.Request) []Artist {

	// fmt.Println(artistsList)
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

		if !FilterFirstAlbum(v, firstAlbum) {
			continue
		}

		//mettre a la fin de la boucle
		if isRight {
			filterArtist = append(filterArtist, v)
		}
	}

	// fmt.Println(filterArtist)

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
