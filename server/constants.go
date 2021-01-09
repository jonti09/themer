package server

import (
	"os"
	"path"
)

var cwd, _ = os.Getwd()
var themesPath = path.Join(cwd, "data", "themes.json")
var themesDir = path.Join(cwd, "data", "themes.json")
var publicDir = path.Join(cwd, "public")
var staticDir = path.Join(publicDir, "static")
