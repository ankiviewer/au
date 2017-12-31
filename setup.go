package main

import (
    "regexp"
    "errors"
    "os/exec"
    "io/ioutil"
    "strings"
)

var args = []Arg{
    Arg{"filestructure", "sets up a great filestructure for the project"},
    Arg{"aliases", "sets up great aliases for changing directory around the project" },
    Arg{"envs", "sets up the required environment variables"},
    // Arg{"remotes", "sets up the heroku remotes"},
}

var SetupCmd = Cmd{
    "sets up aliases for changing directory in the project",
    args,
    setup,
}

var aliases = []struct {
    alias string
    dest string
}{
    {"avweb", "/av_umbrella/apps/av_web"},
    {"avmain", "/av_umbrella/apps/av"},
    {"avassets", "/assets"},
    {"avnodeapp", "/nodeapp"},
}

var projects = []string{
    "av_umbrella",
    "assets",
    "nodeapp"
}

func setup(fp string, arg string) ([]string, error) {
    out := ""
    root := regexp.MustCompile(`.*ankiviewer`).FindString(fp)

    switch arg {
    case "":
        for _, a := range args {
            out += "av setup " + a.name + strings.Repeat(" ", 20 - len(a.name)) + a.desc
        }
    case "filestructure":
        filestructure()
    case "aliases":
    case "envs":
    }

    if root == "" {
        exec.Command("mkdir", "ankiviewer")
        root = fp + "/ankiviewer"
    }
    for _, project := range projects {
        files, _ := ioutil.ReadDir(root)
        if !StringInSlice(project, files) {
            exec.Command(
                "git",
                "clone",
                "https://github.com/ankiviewer/" + project + ".git",
                root + "/" + project
            )
        }
    }

    out += "Configured:" + FileStructure

    aliasAcc := ""
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

eval $(au setup)
`}, nil
    } else {
        return , errors.New("arg not found")
    }
}

