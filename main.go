package main

import (
	"net/http"

	handlers "./src/handlers"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))) // récupère tous les fichiers "externe" dans "static"

	http.HandleFunc("/", handlers.Global)
	http.HandleFunc("/locations", handlers.Locations)
	http.HandleFunc("/artists/", handlers.Artists)

	http.HandleFunc("/search", handlers.Search)

	http.ListenAndServe(":8080", nil)
}
