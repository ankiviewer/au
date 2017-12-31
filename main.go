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

func getEnvs() map[string]string{
    envs := make(map[string]string)
    for _, e := range []string{"AV_ANKI_SQLITE_PATH", "AV_ANKIVIEWER_PATH"} {
        envs[e] = os.Getenv(e)
    }
    return envs
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
    envs := getEnvs()

    o := Opts{fp, arguments, aliases, envs, root, dirsInRoot}

    if command == "help" {
        handleCmd(Help(o))
    } else if c, ok := AvCmds[command]; ok {
        handleCmd(c.f(o))
    } else {
        fmt.Println("Command not found")
    }
}
