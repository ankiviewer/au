package main

var OpenCmd = Cmd{
    "Open Cmd",
    []Arg{},
    open,
}

func open(_ string, _ string) ([]string, error) {
    return []string{"OPEN"}, nil
}
