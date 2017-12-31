package main

var StartCmd = AvCmd{
    "Start Cmd",
    []AvArg{},
    start,
}

func start(o Opts) ([]Command, error) {
    return []Command{}, nil
}
