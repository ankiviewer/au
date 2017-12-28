package main

var StartCmd = Cmd{
    "Start Cmd",
    []Arg{},
    true,
    start,
}

func start(_ string, _ string) ([]string, error) {
    return []string{"START"}, nil
}
