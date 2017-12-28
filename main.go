package main

// import (
//     "fmt"
//     "flag"
//     "path/filepath"
//     "strings"
//     "os"
//     "os/exec"
//     "log"
//     "regexp"
//     "au/messages"
// )
// 
// var argsErrorMessage = "Incorrect args: %s, see usage by typing: au help"
// var notInAppMessage = "Err: not in anki_viewer_umbrella directory"
// var notInAppRootMessage = "Err: the %s command needs to be run from the anki_viewer_umbrella root"

func main() {
}
//     fp, _ := filepath.Abs("./")
// 
//     if !inAnkiApp(fp) {
//       fmt.Println(notInAppMessage)
//       return
//     }
// 
//     args := map[string][]string{
//       "install": {},
//       "build": []string{"-js", "-css", "-elm", "-static", "-w"},
//       "test": []string{"-w", "-node", "-js", "-elm", "-nw"},
//       "cover": []string{"-html", "-open"},
//       "start": []string{"-prod"},
//       "versions": {},
//       "deploy": {},
//     }
// 
//     flag.Parse()
// 
//     a := flag.Args()
// 
//     if len(a) == 0 || stringInSlice(a[0], []string{"-h", "--help", "help"}) {
//         fmt.Println(messages.Usage)
//         return
//     }
// 
//     if !stringInSlice(a[0], mapKeys(cdArgs)) || !stringInSlice(mapKeys(args)) {
//         fmt.Println(fmt.Sprintf(argsErrorMessage, strings.Join(a)))
//         return
//     }
// 
//     if stringInSlice(a[0], mapKeys(cdArgs)) {
//         cd(fp, a[0], os.Chdir)
//         return
//     }
// 
//     if stringInSlice(a[0], mapKeys(args)) && !inAnkiRoot() {
//       fmt.Println(fmt.Sprintf(notInAppRootMessage), a[0])
//       return
//     }
// 
//     if (len(a) == 1) {
//         execute(a)
//     } else {
//         if stringInSlice(a[1], args[a[0]]) {
//             execute(a)
//         } else {
//           fmt.Println(argsErrorMessage, strings.Join(a))
//         }
//     }
// }
// 
// func panicDefault(args []string) {
//     panic("Reached default execution with command: " + strings.Join(args, " "))
// }
// 
// func install() {
//   fmt.Println("install")
// }
// 
// func build(s string) {
//   fmt.Println("build")
// }
// 
// func test(s string) {
//   fmt.Println("test")
// }
// 
// func cover(s string) {
//   fmt.Println("cover")
// }
// 
// func start(s string) {
//   fmt.Println("start")
// }
// 
// func versions() {
//     standard_v_regex := `(\d\d?\.\d\d?(?:\.\d\d?)?)`
//     elixir_v_regex := `(\d\d?\.\d\d?(?:\.\d\d?)?)(?:\s*)?$`
//     execs := make(map[string][]string)
//     execs["goon"] = append(execs["goon"], "-v", standard_v_regex)
//     execs["postgres"] = append(execs["postgres"], "--version", standard_v_regex)
//     execs["heroku"] = append(execs["heroku"], "--version", standard_v_regex)
//     execs["sass"] = append(execs["sass"], "--version", standard_v_regex)
//     execs["node"] = append(execs["node"], "--version", standard_v_regex)
//     execs["elm"] = append(execs["elm"], "--version", standard_v_regex)
//     execs["elixir"] = append(execs["elixir"], "--version", elixir_v_regex)
// 
//     for k, v := range execs {
//         _, err := exec.LookPath(k)
//         if err != nil {
//             log.Fatal(k + " executable not in your $PATH")
//         }
//         out, err := exec.Command(k, v[0]).Output()
//         if err != nil {
//             log.Fatal(err)
//         }
//         re := regexp.MustCompile(v[1])
//         fmt.Println(k + ": v" + re.FindStringSubmatch(string(out))[1])
//     }
// }
// 
// func execute(args []string) {
//     switch args[0] {
//     case "install":
//         install()
//     case "build":
//         if len(args) == 1 {
//             build("")
//         } else if len(args) == 2 {
//             build(args[1])
//         } else if args[2] != "-w" {
//             build("fail")
//         } else {
//             build(args[1] + ":w")
//         }
//     case "test":
//         if len(args) == 1 {
//             test("")
//         } else if len(args) == 2 {
//             test(args[1])
//         } else if args[2] != "-w" {
//             test("fail")
//         } else {
//             test(args[1] + ":w")
//         }
//     case "cover":
//         if len(args) == 1 {
//             cover("")
//         } else {
//             cover(args[1])
//         }
//     case "start":
//         if len(args) == 1 {
//             start("")
//         } else {
//             start(args[1])
//         }
//     case "versions":
//         versions()
//     default:
//         panicDefault(args)
//     }
// }
