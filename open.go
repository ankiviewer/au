package main

var OpenCmd = Cmd{
    "Open Cmd",
    []Arg{Arg{"travis", "opens travis"}},
    open,
}

func open(_ string, _ string) ([]string, error) {
    return []string{"OPEN"}, nil
}
