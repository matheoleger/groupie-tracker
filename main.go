package main

import (
	"html/template"
	"net/http"
)

var t *template.Template

func main() {
	// Récupération du fichier HTML pour la mise en forme
	t = template.Must(template.ParseFiles("./index.html"))

	// Prise en compte les fichiers dans le dossier static (CSS et fichiers .txt)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//Exécution de la fonction home lorsque le serveur est ouvert
	//http.HandleFunc("/", home)

	//Exécution de la fonction print lorsque le client ira sur la page /ascii-art
	//Paramétré dans le formulaire (index.html) avec action="/ascii-art L28"
	//http.HandleFunc("/ascii-art", print)

	//Définition du port utilisé pour ce serveur
	http.ListenAndServe(":8000", nil)
}
