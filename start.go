package main

var StartCmd = Cmd{
    "Start Cmd",
    []Arg{},
    start,
}

func start(_ string, _ string) ([]string, error) {
    return []string{"START"}, nil
}
