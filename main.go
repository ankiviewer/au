package main

import (
    "os"
    "path/filepath"
    "fmt"
)

type Arg struct {
    name string
    desc string
}

type Cmd struct {
    desc string
    args []Arg
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

func handleLog(ss []string, err error) {
    if err != nil {
        panic(err)
    } else {
        for _, s := range ss {
            fmt.Println(s)
        }
    }
}

func main() {
    fp, _ := filepath.Abs("./")
    arguments := os.Args[1:]
    command := arguments[0]

    if c, ok := AuCmds[command]; ok {
        switch {
        case len(arguments) == 0:
            handleLog(c.f(fp, ""))
        case len(arguments) == 1:
            handleLog(c.f(fp, arguments[1]))
        default:
            // Consider allowing more than 1 argument in the future
            panic("Too many arguments")
        }
    } else {
        panic("Command not found")
    }
}
