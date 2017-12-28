package main

var TestCmd = Cmd{
    "Test Cmd",
    []Arg{},
    true,
    test,
}

func test(fp string, _ string) ([]string, error) {
    return []string{"TEST"}, nil
}
