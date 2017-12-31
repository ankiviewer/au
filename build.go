package main

var BuildCmd = AvCmd{
    "Build Cmd",
    []AvArg{},
    build,
}

func build(o Opts) ([]Command, error) {
    return []Command{}, nil
}
