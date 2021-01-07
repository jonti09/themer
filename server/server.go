package server

import (
	"fmt"
	"net/http"
)

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	themes, err := GetListOfThemes()

	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}

	_, _ = w.Write(themes)
}

// StartAndListen method starts the server on 8000 port
func StartAndListen() {
	http.HandleFunc("/themes/", homePage)
	fmt.Printf("Starting the server on %d...\n", 8000)
	http.ListenAndServe(":8000", nil)

}
