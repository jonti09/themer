package server

import (
	"io/ioutil"
	"os"
	"path"
)

func getListOfThemes() ([]byte, error) {
	themesFile, err := os.Open(path.Join("data", "themes.json"))
	defer themesFile.Close()

	themesBytes, err := ioutil.ReadAll(themesFile)

	if err != nil {
		return nil, err
	}

	return themesBytes, nil
}
