package main

var MessagesCmd = Cmd{
    "Messages Cmd",
    []Arg{},
    true,
    messages,
}

func messages(_ string, _ string) ([]string, error) {
    return []string{"MESSAGES"}, nil
}
