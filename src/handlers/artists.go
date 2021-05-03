package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
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

	} else {
		http.Error(w, "Page not found", 404)
	}
	// ts, err := template.ParseFiles("./templates/layout.html", "./templates/location.html")

}
