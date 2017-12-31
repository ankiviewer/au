package main

var TestCmd = AvCmd{
    "Test Cmd",
    []AvArg{},
    test,
}

func test(o Opts) ([]Command, error) {
    return []Command{}, nil
}
