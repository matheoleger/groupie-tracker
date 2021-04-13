package main

import (
	"net/http"

	handlers "./src/handlers"
)

func main() {

	// http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css")))) // récupère tous les fichiers "externe" dans "css"
	// http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./static/img")))) // récupère tous les fichiers "externe" dans "images"
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/location", handlers.Location)
	http.ListenAndServe(":8080", nil)
}
