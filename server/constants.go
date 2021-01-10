package server

import (
	"os/user"
	"path"
)

var asset = "public.tgz"

var usr, _ = user.Current()
var homeDir = usr.HomeDir
var workingDir = path.Join(homeDir, ".config", "themer")

var publicDir = path.Join(workingDir, "public")
var themesPath = path.Join(publicDir, "data", "themes.json")
var themesDir = path.Join(publicDir, "data", "themes")
var staticDir = path.Join(publicDir, "static")
