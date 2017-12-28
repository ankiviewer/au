package main

var OpenCmd = Cmd{
    "Open Cmd",
    []Arg{},
    true,
    open,
}

func open(_ string, _ string) ([]string, error) {
    return []string{"OPEN"}, nil
}
