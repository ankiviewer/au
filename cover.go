package main

var CoverCmd = AvCmd{
    "Cover Cmd",
    []AvArg{},
    cover,
}

func cover(o Opts) ([]Command, error) {
    return []Command{}, nil
}
