package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func sendResponseOrError(w http.ResponseWriter, err error, data []byte) {
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write(data)

	if err != nil {
		log.Printf("There was an error in writing the response: %s", err)
	}
}

func listThemes(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)

	themes, err := GetListOfThemes()
	sendResponseOrError(w, err, themes)
}

type themeStruct struct {
	Theme string
}

func applyTheme(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "POST" {
		var theme themeStruct
		err := json.NewDecoder(r.Body).Decode(&theme)

		if err != nil {
			sendResponseOrError(w, err, []byte(""))
			return
		}

		_, err = ApplyTheme(theme.Theme)
		sendResponseOrError(w, err, []byte("Theme applied successfully"))
	}
}

// StartAndListen method starts the server on 8000 port
func StartAndListen() {
	http.HandleFunc("/themes/", listThemes)
	http.HandleFunc("/apply/", applyTheme)
	log.Printf("Starting the server on %d...\n", 8000)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
