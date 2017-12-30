package main

import (
  "strings"
  "fmt"
  "github.com/ankiviewer/av/messages"
)

func spaces(n int) string {
  return strings.Repeat(" ", n)
}

func formatCmd(name string, desc string) string {
  return "av " + name + spaces(20 - len(name)) + desc + "\n"
}

func formatArg(cname string, name string, desc string) string {
  return "av " + cname + " " + name + spaces(15 - len(name)) + desc + "\n"
}

func Help(_ string, _ string) ([]string, error) {
    out := formatCmd("help", "shows this help prompt\n")
    for k, v := range AvCmds {
        out += formatCmd(k, v.desc)
        argOut := ""
        for _, a := range v.args {
            argOut += formatArg(k, a.name, a.desc)
        }
        out += argOut
        out += "\n"
    }

    out = fmt.Sprintf(messages.Usage, out)
    return []string{out}, nil
}
