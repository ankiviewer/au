package main

import (
    "os/exec"
    "regexp"
    "errors"
    "strings"
)

var VersionsCmd = AvCmd{
    "prints the versions of all the project technologies",
    []AvArg{},
    versions,
}

// TODO: ensure versions are met
type executable struct {
    name string
    versionFlag string
    versionRegex string
}

var standard_v_regex = `(\d\d?\.\d\d?(?:\.\d\d?)?)`
var elixir_v_regex = `(\d\d?\.\d\d?(?:\.\d\d?)?)(?:\s*)?$`

var shortV = "-v"
var longV = "--version"

var executables = []executable{
    executable{"goon", shortV, standard_v_regex},
    executable{"postgres", longV, standard_v_regex},
    executable{"heroku", longV, standard_v_regex},
    executable{"sass", longV, standard_v_regex},
    executable{"node", longV, standard_v_regex},
    executable{"elm", longV, standard_v_regex},
    executable{"elixir", longV, elixir_v_regex},
}

func versions(o Opts) ([]Command, error) {
    outs := make([]string, 0, len(executables))
    var err error

    for _, ex := range executables {
        _, e := exec.LookPath(ex.name)
        if e != nil {
            err = errors.New(ex.name + " executable not in $PATH")
            break
        }
        out, e := exec.Command(ex.name, ex.versionFlag).Output()
        if e != nil {
            err = e
            break
        }
        re := regexp.MustCompile(ex.versionRegex)
        outs = append(outs, ex.name + ": v" + re.FindStringSubmatch(string(out))[1])
    }

    if err != nil {
        return []Command{}, err
    } else {
        return []Command{Command{strings.Join(outs, "\n"), ""}}, nil
    }
}
