package main

var InstallCmd = Cmd{
    "Install Cmd",
    []Arg{},
    install,
}

func install(_ string, _ string) ([]string, error) {
    return []string{"INSTALL"}, nil
}
