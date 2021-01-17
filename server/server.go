package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type themeStruct struct {
	Theme string
}

func listThemesAPI(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)

	themes, err := getListOfThemes()
	sendResponseOrError(w, err, themes)
}

func applyThemeAPI(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	var output []byte

	if r.Method == "POST" {
		var theme themeStruct
		err := json.NewDecoder(r.Body).Decode(&theme)

		if err != nil {
			sendResponseOrError(w, err, []byte(""))
			return
		}

		output, err = applyTheme(theme.Theme)
		if err != nil {
			// print error for debugging
			log.Printf(string(output))
		}
		sendResponseOrError(w, err, []byte("Theme applied successfully"))
	}
}

func setupStaticHandler() {
	http.Handle("/", http.FileServer(http.Dir(publicDir)))
	http.Handle("/static/", http.FileServer(http.Dir(publicDir)))
}

// StartAndListen method starts the server on 8000 port
func StartAndListen() {
	setupStaticHandler()
	http.HandleFunc("/themes/", listThemesAPI)
	http.HandleFunc("/apply/", applyThemeAPI)
	log.Printf("Starting the server on http://localhost:%d/\n", 8000)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
