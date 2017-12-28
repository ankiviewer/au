package main

var TestCmd = Cmd{
    "Test Cmd",
    []Arg{},
    test,
}

func test(fp string, _ string) ([]string, error) {
    return []string{"TEST"}, nil
}
