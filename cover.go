package main

var CoverCmd = Cmd{
    "Cover Cmd",
    []Arg{},
    true,
    cover,
}

func cover(_ string, _ string) ([]string, error) {
    return []string{"COVER"}, nil
}
