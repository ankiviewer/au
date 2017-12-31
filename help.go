package main

// import (
//   "strings"
//   "fmt"
//   "github.com/ankiviewer/av/messages"
// )
// 
// func formatCmd(name string, desc string) string {
//   return "av " + name + Spaces(20 - len(name)) + desc + "\n"
// }
// 
// func formatArg(cname string, name string, desc string) string {
//   return "av " + cname + " " + name + Spaces(15 - len(name)) + desc + "\n"
// }

func Help(o Opts) ([]Command, error) {
    return []Command{}, nil
    // out := formatCmd("help", "shows this help prompt\n")
    // for k, v := range AvCmds {
    //     out += formatCmd(k, v.desc)
    //     argOut := ""
    //     for _, a := range v.args {
    //         argOut += formatArg(k, a.name, a.desc)
    //     }
    //     out += argOut
    //     out += "\n"
    // }

    // out = fmt.Sprintf(messages.Usage, out)
    // return []string{out}, nil
}
