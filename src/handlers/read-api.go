/*package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)



type band struct {
	ID int `json:"id"`
	/*Image struct {
		Xs string
	}
	Name string `json:"name"`
}

type PAgeDataArtist struct {
	artists []band
}

func DEfaultPage(w http.ResponseWriter, r *http.Request) {
	var numberOfArtist int = 20
	var firstID int = 1
	var firstIDNewPAge int
	fmt.Print(APIArtist)

	if r.URL.Path != "/artists" {
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
		} else {
			firstID = firstID + firstIDNewPAge
		}

	}
}

func GetArtist(firstID int, numArtist int) []band {
	var ArtistList []band = []band{}

	if numArtist == 0 {
		res, err := http.Get(ALLJSONURL)
		if err != nil {
			fmt.Print("NO API", err)
			return (nil)
		}

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("CAN't READ")
			return (nil)
		}
		json.Unmarshal(data, ArtistList)
	} else {
		i := firstID
		for i < firstID+numArtist {
			res, err := http.Get(APIArtist + "id/" + strconv.Itoa(i) + ".json")
			if err != nil {
				fmt.Println("Dindn't get the info", err)
				return (nil)
			}
			data, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println("Can'i read", err)
				return (nil)
			}

			var tmpArtist band
			json.Unmarshal(data, &tmpArtist)
			if tmpArtist.ID != 0 {
				fmt.Println(i)
				ArtistList = append(ArtistList, tmpArtist)
				i++
			}
			// DEBUG HERE
		}
	}
	fmt.Println(ArtistList)
	return ArtistList
}

func main() {
	GetArtist(1, 200)
}
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const APIArtist string = "https://groupietrackers.herokuapp.com/api/"

const ALLJSONURL string = "https://groupietrackers.herokuapp.com/api/artists/all.json"

func main() {
	/*response, err := http.Get(APIArtist)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}*/
	response := getAPI("artists")

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

}

func getAPI(s string) *http.Response {
	reponse, err := http.Get(APIArtist + s)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(-1)
	}
	return (reponse)
}
