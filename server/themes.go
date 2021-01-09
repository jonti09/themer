package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
)

// GetListOfThemes return list of themes available in `themes.json` file
func GetListOfThemes() ([]byte, error) {
	themesFile, err := os.Open(path.Join("data", "themes.json"))
	defer themesFile.Close()

	themesBytes, err := ioutil.ReadAll(themesFile)

	if err != nil {
		return nil, err
	}

	return themesBytes, nil
}

// ApplyTheme function takes a theme name and installs it in current terminal
func ApplyTheme(theme string) ([]byte, error) {
	log.Println("Applying the theme " + theme + " ...")
	cwd, _ := os.Getwd()
	themeFilePath := path.Join(cwd, "data", "themes", theme)
	cmd := exec.Command("x-terminal-emulator", "-e", "bash", themeFilePath)

	fmt.Println(cmd)

	return cmd.CombinedOutput()
}
