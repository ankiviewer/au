package main

var BuildCmd = Cmd{
    "Build Cmd",
    []Arg{},
    true,
    build,
}

func build(_ string, _ string) ([]string, error) {
    return []string{"BUILD"}, nil
}
