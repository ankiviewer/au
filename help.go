package main

var HelpCmd = Cmd{
    "Help Cmd",
    []Arg{},
    true,
    help,
}

func help(_ string, _ string) ([]string, error) {
    return []string{"HELP"}, nil
}
