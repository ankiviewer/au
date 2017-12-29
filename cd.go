package main

import (
    "regexp"
    "errors"
)

var CdCmd = Cmd{
    "sets up aliases for changing directory in the project",
    []Arg{},
    cd,
}

var aliases = []struct {
    alias string
    dest string
}{
    {"auroot", ""},
    {"auweb", "/apps/anki_web"},
    {"auassets", "/apps/anki_web/assets"},
    {"aumain", "/apps/anki"},
    {"aunodeapp", "/apps/anki/nodeapp"},
}

func cd(fp string, arg string) ([]string, error) {
    root := regexp.MustCompile(`.*anki_viewer_umbrella`).FindString(fp)
    aliasAcc := ""
    if arg == "output" {
        for _, a := range aliases {
            x := "alias " + a.alias + "=" + root + a.dest + ";"
            aliasAcc = aliasAcc + x
        }

        return []string{aliasAcc}, nil
    } else if arg == "" {
        return []string{`To set up the useful aliases:

auroot
auweb
auassets
aumain
aunodeapp

Run the following command in your terminal:

eval $(au cd)
`}, nil
    } else {
        return nil, errors.New("arg not found")
    }
}

