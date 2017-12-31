package main

var OpenCmd = AvCmd{
    "Open Cmd",
    []AvArg{AvArg{"travis", "opens travis"}},
    open,
}

func open(o Opts) ([]Command, error) {
    return []Command{}, nil
}
