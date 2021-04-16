package handlers

import (
	"fmt"
	"net/http"
)

const APIArtist string = "https://groupietrackers.herokuapp.com/api/artists"

//const ALLJSONURL string = "https://groupietrackers.herokuapp.com/api/all.json"

type band struct {
	ID    int
	Image struct {
		Xs string
	}
	Name string
}

func DEfaultPage(w http.ResponseWriter, r *http.Request) {
	var numberOfArtist int = 20
	var firstID int = 1
	fmt.Print(APIArtist)

	if r.URL.Path != "/information" {
		http.NotFound(w, r)
		return
	}

	if r.Method == "POST" {
		if r.FormValue(("numOfArtist")) != "" {
			switch r.FormValue("numOfArtist") {
			case "50":
				numberOfArtist = 50
				break
			case "100":
				numberOfArtist = 100
				break
			case "all":
				numberOfArtist = 0
				break
			default:
				numberOfArtist = 20
			}
		}
		if numberOfArtist == 0 {
			firstID = 1
		}

	}
}
