package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
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

// GetListOfThemes return list of themes available in `themes.json` file
func getListOfThemes() ([]byte, error) {
	themesFile, err := os.Open(themesPath)
	defer themesFile.Close()

	themesBytes, err := ioutil.ReadAll(themesFile)

	if err != nil {
		return nil, err
	}

	return themesBytes, nil
}

// ApplyTheme function takes a theme name and installs it in current terminal
func applyTheme(theme string) ([]byte, error) {
	terminal := os.Getenv("TERMINAL")
	log.Println("Applying the theme: " + theme + " in terminal " + terminal)
	themeFilePath := path.Join(themesDir, theme)
	cmd := exec.Command(themeFilePath)
	return cmd.CombinedOutput()
}

func fileExists(name string) bool {
	info, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
