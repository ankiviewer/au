package main

import (
  "regexp"
)

var CdArgs = map[string]string{
  "root": "",
  "..": "",
  "web": "apps/anki_web",
  "assets": "/apps/anki_web/assets",
  "anki": "/apps/anki",
  ".": "/apps/anki",
  "nodeapp": "/apps/anki/nodeapp",
}

func CdDir(currentFilePath string, arg string) string {
    root := regexp.MustCompile(`.*anki_viewer_umbrella`).FindString(currentFilePath)

    desiredDir := root + CdArgs[arg]

    return desiredDir
}

