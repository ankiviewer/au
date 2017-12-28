package main

import (
    "os"
    "path/filepath"
    "fmt"
    "strings"
)

type Arg struct {
    name string
    desc string
}

type Cmd struct {
    desc string
    args []Arg
    log bool
    f func(string, string) ([]string, error)
}

var AuCmds = map[string]Cmd{
    "help": HelpCmd,
    "cd": CdCmd,
    "open": OpenCmd,
    "install": InstallCmd,
    "build": BuildCmd,
    "test": TestCmd,
    "cover": CoverCmd,
    "start": StartCmd,
    "versions": VersionsCmd,
    "deploy": DeployCmd,
}

func handleLog(s []string, _ error) {
  fmt.Println(strings.Join(s, ""))
}

func main() {
    fp, _ := filepath.Abs("./")
    arguments := os.Args[1:]
    command := arguments[0]

    if c, ok := AuCmds[command]; ok {
        switch {
        case len(arguments) == 0 && c.log:
            handleLog(c.f(fp, ""))
        case len(arguments) == 0 && !c.log:
            c.f(fp, "")
        case len(arguments) == 1 && c.log:
            handleLog(c.f(fp, arguments[1]))
        case len(arguments) == 1 && !c.log:
            c.f(fp, arguments[1])
        default:
            // Consider allowing more than 1 argument in the future
            panic("Too many arguments")
        }
    } else {
        panic("Command not found")
    }
}
