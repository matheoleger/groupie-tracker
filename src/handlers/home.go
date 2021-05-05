package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func Global(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {

		//ParseFiles très important : on doit y mettre toutes les pages qui doivent être chargé et donc choisir en fonction de celle qui doit l'être
		// le choix doit être fait avant le ParseFiles
		files := findPathFiles("./templates/index.html")
		ts, err := template.ParseFiles(files...)

		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}

		ts.Execute(w, nil)
	} else {
		http.Error(w, "Page not found", 404)
	}

}
