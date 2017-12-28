package main

var DeployCmd = Cmd{
    "Deploy Cmd",
    []Arg{},
    true,
    deploy,
}

func deploy(_ string, _ string) ([]string, error) {
    return []string{"DEPLOY"}, nil
}
