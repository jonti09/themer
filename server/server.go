package server

import (
	"fmt"
	"log"
	"net/http"
)

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	themes, err := getListOfThemes()

	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}

	_, _ = w.Write(themes)
}

func StartAndListen() {
	http.HandleFunc("/themes/", homePage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
