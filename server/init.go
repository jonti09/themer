package server

import (
	"log"
	"os"
	"os/exec"
)

func makeAndChangeDir() {
	log.Println("Creating and changing directory")
	err := os.MkdirAll(workingDir, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	err = os.Chdir(workingDir)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Changed the working directory to: " + workingDir)
}

func downloadAssets() {
	log.Println("Downloading necessary assets")
	cmd := exec.Command("wget", "--content-disposition", "https://github.com/jonti09/themer/releases/download/latest/"+asset, "-O", asset)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("%s error with %s", err, output)
	}
	log.Println("Asset downloaded")
}

func setupAssets() {
	log.Println("Extracting the assets")
	cmd := exec.Command("tar", "-xvzf", asset)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("%s error with %s", err, output)
	}
	log.Println("Assets extraction complete")
}

// Init will initialize the working directory, get needed assets from Git
func Init() {
	makeAndChangeDir()

	// only download assets if it does not exists
	if !fileExists(asset) {
		downloadAssets()
	} else {
		log.Println("Asset exists! Skipping download")
	}

	setupAssets()
}
