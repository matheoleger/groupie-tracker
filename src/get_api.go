package usefulFiles

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetApi(path string) []byte {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/" + path)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}
