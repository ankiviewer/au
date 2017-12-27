package main

import (
    "fmt"
    "flag"
    "path/filepath"
    "strings"
    "os"
    "os/exec"
    "log"
    "regexp"
    "au/messages"
)

var argsErrorMessage = `Incorrect args, see usage by typing: au help`

func inAnkiRoot() bool {
    fp, _ := filepath.Abs("./")

    return strings.HasSuffix(fp, "anki_viewer_umbrella")
}

func inAnkiApp() bool {
    fp, _ := filepath.Abs("./")

    return strings.Contains(fp, "anki_viewer_umbrella")
}

func getAnkiRoot() string {
    fp, _ := filepath.Abs("./")

    return regexp.MustCompile(`.*anki_viewer_umbrella`).FindString(fp)
}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func cd(dir string) {
    switch dir {
    case "root":
        os.Chdir(getAnkiRoot())
    case "..":
        os.Chdir(getAnkiRoot())
    case "web":
        os.Chdir(getAnkiRoot() + "/apps/anki_web")
    case "assets":
        os.Chdir(getAnkiRoot() + "/apps/anki_web/assets")
    case "anki":
        os.Chdir(getAnkiRoot() + "/apps/anki")
    case ".":
        os.Chdir(getAnkiRoot() + "/apps/anki")
    case "nodeapp":
        os.Chdir(getAnkiRoot() + "/apps/anki/nodeapp")
    default:
        panic("Shouldn't reach here, arg used: " + dir)
    }
}

func main() {
    cdArgs := []string{"root", "..", "web", "assets", "anki", ".", "nodeapp"}
    firstArgs := []string{"install", "build", "test", "cover", "start", "versions", "deploy"}
    secondArgs := make(map[string][]string)

    secondArgs["build"] = append(secondArgs["build"], "-js", "-css", "-elm", "-static", "-w")
    secondArgs["test"] = append(secondArgs["test"], "-w", "-node", "-js", "-elm", "-nw")
    secondArgs["cover"] = append(secondArgs["cover"], "-html", "-open")
    secondArgs["start"] = append(secondArgs["start"], "-prod")

    flag.Parse()

    a := flag.Args()

    if len(a) == 0 || a[0] == "-h" || a[0] == "--help" || a[0] == "help" {
          fmt.Println(messages.Usage)
          return
    }

    if !stringInSlice(a[0], append(cdArgs, firstArgs...)) {
        panic("Unknown command: " + a[0])
    }

    if stringInSlice(a[0], firstArgs) && !inAnkiRoot() {
        panic("Run the " + a[0] + " au command from the root of the umbrella project")
    }

    if stringInSlice(a[0], cdArgs) {
      cd(a[0])
    }

    if flags, exists := secondArgs[a[0]]; exists {
        if (len(a) == 1) {
            execute(a)
        } else {
            if stringInSlice(a[1], flags) {
                execute(a)
            } else {
                fmt.Println(argsErrorMessage)
            }
        }
    } else {
        fmt.Println(argsErrorMessage)
    }
}

func panicDefault(args []string) {
    panic("Reached default execution with command: " + strings.Join(args, " "))
}

func install() {
  fmt.Println("install")
}

func build(s string) {
  fmt.Println("build")
}

func test(s string) {
  fmt.Println("test")
}

func cover(s string) {
  fmt.Println("cover")
}

func start(s string) {
  fmt.Println("start")
}

func versions() {
    standard_v_regex := `(\d\d?\.\d\d?(?:\.\d\d?)?)`
    elixir_v_regex := `(\d\d?\.\d\d?(?:\.\d\d?)?)(?:\s*)?$`
    execs := make(map[string][]string)
    execs["goon"] = append(execs["goon"], "-v", standard_v_regex)
    execs["postgres"] = append(execs["postgres"], "--version", standard_v_regex)
    execs["heroku"] = append(execs["heroku"], "--version", standard_v_regex)
    execs["sass"] = append(execs["sass"], "--version", standard_v_regex)
    execs["node"] = append(execs["node"], "--version", standard_v_regex)
    execs["elm"] = append(execs["elm"], "--version", standard_v_regex)
    execs["elixir"] = append(execs["elixir"], "--version", elixir_v_regex)

    for k, v := range execs {
        _, err := exec.LookPath(k)
        if err != nil {
            log.Fatal(k + " executable not in your $PATH")
        }
        out, err := exec.Command(k, v[0]).Output()
        if err != nil {
            log.Fatal(err)
        }
        re := regexp.MustCompile(v[1])
        fmt.Println(k + ": v" + re.FindStringSubmatch(string(out))[1])
    }
}

func execute(args []string) {
    switch args[0] {
    case "install":
        install()
    case "build":
        if len(args) == 1 {
            build("")
        } else if len(args) == 2 {
            build(args[1])
        } else if args[2] != "-w" {
            build("fail")
        } else {
            build(args[1] + ":w")
        }
    case "test":
        if len(args) == 1 {
            test("")
        } else if len(args) == 2 {
            test(args[1])
        } else if args[2] != "-w" {
            test("fail")
        } else {
            test(args[1] + ":w")
        }
    case "cover":
        if len(args) == 1 {
            cover("")
        } else {
            cover(args[1])
        }
    case "start":
        if len(args) == 1 {
            start("")
        } else {
            start(args[1])
        }
    case "versions":
        versions()
    default:
        panicDefault(args)
    }
}
