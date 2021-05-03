package handlers

import (
	"fmt"
	"net/http"
)

func howMany(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Query().Get("filtre"))
}
