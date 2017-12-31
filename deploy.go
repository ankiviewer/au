package main

var DeployCmd = AvCmd{
    "Deploy Cmd",
    []AvArg{},
    deploy,
}

func deploy(o Opts) ([]Command, error) {
    return []Command{}, nil
}
