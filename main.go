//+build !test

package main

import (
    "os"
    "os/exec"
    "path/filepath"
    "fmt"
    "regexp"
    "strings"
    "github.com/ankiviewer/av/messages"
)

func handleCmd(cmds []Command, err error) {
    if err != nil {
        fmt.Println(err)
    } else {
        for _, c := range cmds {
            if c.cmd != "" {
                _, cerr := exec.Command("sh", "-c", c.cmd).Output()
                if cerr != nil {
                    fmt.Printf("CMD ERR: %s", cerr)
                    break
                }
            }
            if c.log != "" {
                fmt.Println(c.log)
            }
        }
    }
}

func getAliases() map[string]string{
    aliases := make(map[string]string)
    for _, a := range []string{"avweb", "avmain", "avassets", "avnodeapp"} {
        aout, aerr := exec.Command("alias", a).Output()
        if aerr != nil {
            aliases[a] = ""
        } else {
            aliases[a] = string(aout)
        }
    }
    return aliases
}

func getSqlitePath() string {
    c := "find \"$(find $HOME -type d | grep Anki2 | head -1)\" -type f | grep collection\\.anki2$"
    out, err := exec.Command("sh", "-c", c).Output()
    if err != nil {
        return ""
    }
    return strings.Trim(string(out), "\n ")
}

func main() {
    if len(os.Args) == 1 {
        fmt.Println(messages.NoInput)
        return
    }

    fp, _ := filepath.Abs("./")
    root := regexp.MustCompile(`.*ankiviewer`).FindString(fp)
    dirsInRoot := []string{}
    if root != "" {
        lsInRoot, _ := exec.Command("ls", root).Output()
        for _, dir := range strings.Split(string(lsInRoot), "\n") {
            dirsInRoot = append(dirsInRoot, dir)
        }
    }
    command := os.Args[1]
    arguments := os.Args[2:]
    aliases := getAliases()
    sqlitePath := os.Getenv("AV_ANKI_SQLITE_PATH")
    ankiviewerPath := os.Getenv("AV_ANKIVIEWER_PATH")
    if sqlitePath == "" {
        sqlitePath = getSqlitePath()
    }

    o := Opts{fp, arguments, aliases, root, dirsInRoot, sqlitePath, ankiviewerPath}

    if command == "help" {
        handleCmd(Help(o))
    } else if c, ok := AvCmds[command]; ok {
        handleCmd(c.f(o))
    } else {
        fmt.Println("Command not found")
    }
}
