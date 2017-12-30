package main

import (
    "os"
    "path/filepath"
    "fmt"
    "github.com/ankiviewer/av/messages"
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

var AvCmds = map[string]Cmd{
    "setup": SetupCmd,
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
        fmt.Println(err)
    } else {
        for _, s := range ss {
            fmt.Println(s)
        }
    }
}

func main() {
    if len(os.Args) == 1 {
        fmt.Println(messages.NoInput)
        return
    }

    fp, _ := filepath.Abs("./")
    command := os.Args[1]
    arguments := os.Args[2:]

    if command == "help" {
      handleLog(Help(fp, ""))
      return
    }

    if c, ok := AvCmds[command]; ok {
        switch {
        case len(arguments) == 0:
            handleLog(c.f(fp, ""))
        case len(arguments) == 1:
            handleLog(c.f(fp, arguments[0]))
        default:
            // Consider allowing more than 1 argument in the future
            panic("Too many arguments")
        }
    } else {
        fmt.Println("Command not found")
    }
}
