package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func Locations(w http.ResponseWriter, r *http.Request) {

	files := findPathFiles("./templates/locations.html")

	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	ts.Execute(w, nil)
}
