package main

import (
  "regexp"
  "os"
  "errors"
)

var CdCmd = Cmd{
  "Changes directory inside project",
  []Arg{},
  false,
  cd,
}

var cdMap = map[string]string{
  "root": "",
  "..": "",
  "web": "apps/anki_web",
  "assets": "/apps/anki_web/assets",
  "anki": "/apps/anki",
  ".": "/apps/anki",
  "nodeapp": "/apps/anki/nodeapp",
}

func cd(fp string, arg string) ([]string, error) {
    root := regexp.MustCompile(`.*anki_viewer_umbrella`).FindString(fp)
    if dest, ok := cdMap[arg]; ok {
        os.Chdir(root + dest)
        return []string{}, nil
    } else {
        return nil, errors.New("arg not found")
    }
}

