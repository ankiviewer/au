package main

var InstallCmd = AvCmd{
    "Install Cmd",
    []AvArg{},
    install,
}

func install(o Opts) ([]Command, error) {
    return []Command{}, nil
}
