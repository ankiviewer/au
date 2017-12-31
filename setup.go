package main

import (
    // "regexp"
    "errors"
    // "os/exec"
    // "io/ioutil"
    // "strings"
)

var args = []AvArg{
    AvArg{
      "filestructure",
      "sets up a great filestructure for the project",
    },
    AvArg{
      "aliases",
      "sets up great aliases for changing directory around the project",
    },
    AvArg{
      "envs",
      "sets up the required environment variables",
    },
    // Arg{"remotes", "sets up the heroku remotes"},
}

var SetupCmd = AvCmd{
    "sets up aliases for changing directory in the project",
    args,
    setup,
}

var projects = []string{
    "av_umbrella",
    "assets",
    "nodeapp",
}

func setup(o Opts) ([]Command, error) {
    a := ""
    if len(o.args) != 0 {
        a = o.args[0]
    }
    switch a {
    case "":
        out := []Command{Command{"Usage:", ""}}

        for _, a := range args {
            out = append(
                out,
                Command{
                    "av setup " + a.name + Spaces(20 - len(a.name)) + a.desc,
                    "",
                },
            )
        }
        return out, nil
    case "filestructure":
        return filestructure(o)
    case "aliases":
        return aliases(o)
    case "envs":
        return envs(o)
    default:
      return []Command{}, errors.New("Command not found: av setup " + a)
    }
}

func filestructure(o Opts) ([]Command, error) {
    switch {
    case o.root == "":
        out := []Command{}
        out = append(
            out,
            Command{
                "making directory ankiviewer",
                "mkdir " + o.fp + "/ankiviewer",
            },
        )
        for _, project := range projects {
            out = append(
                out,
                Command{
                    "cloning " + project,
                    "git clone https://github.com/ankiviewer/" + project + ".git " + o.root + "/" + project,
                },
            )
        }
        return out, nil
    default:
        out := []Command{}
        for _, project := range projects {
            if !StringInSlice(project, o.dirsInRoot) {
                out = append(
                    out,
                    Command{
                        "cloning " + project,
                        "git clone https://github.com/ankiviewer/" + project + ".git " + o.root + "/" + project,
                    },
                )
            }
        }
        out = append(out, Command{"All up to date!", ""})
        return out, nil
    }
}

func aliases(o Opts) ([]Command, error) {
    return []Command{Command{"Hello setup aliases", ""}}, nil
}

func envs(o Opts) ([]Command, error) {
    return []Command{Command{"Hello setup envs", ""}}, nil
}
