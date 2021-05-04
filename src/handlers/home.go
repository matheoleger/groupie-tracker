package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func Global(w http.ResponseWriter, r *http.Request) {

	//ParseFiles très important : on doit y mettre toutes les pages qui doivent être chargé et donc choisir en fonction de celle qui doit l'être
	// le choix doit être fait avant le ParseFiles

	if r.URL.Path == "/" {
		files := findPathFiles("./templates/index.html")

		// ts, err := template.ParseFiles("./templates/layout.html", "./templates/index.html")
		ts, err := template.ParseFiles(files...)

		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}

		ts.Execute(w, nil)
	} else if r.URL.Path == "/search" {

	} else {
		GetApi(w, r)
	}

}
