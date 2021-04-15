package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func Location(w http.ResponseWriter, r *http.Request) {

	//ParseFiles très important : on doit y mettre toutes les pages qui doivent être chargé et donc choisir en fonction de celle qui doit l'être
	// le choix doit être fait avant le ParseFiles
	files := findPathFiles("./templates/locations.html")

	// ts, err := template.ParseFiles("./templates/layout.html", "./templates/location.html")
	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	ts.Execute(w, nil)
}
