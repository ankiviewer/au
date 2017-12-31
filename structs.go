package main

type AvArg struct {
  name string
  desc string
}

type AvCmd struct {
    desc string
    args []AvArg
    f func(Opts) ([]Command, error)
}

type Opts struct {
    fp string
    args []string
    aliases map[string]string
    envs map[string]string
    root string
    dirsInRoot []string
}

type Command struct {
    log string
    cmd string
}

var AvCmds = map[string]AvCmd{
    // can't include "help" command as it references this slice
    "setup": SetupCmd,
    // "open": OpenCmd,
    // "install": InstallCmd,
    // "build": BuildCmd,
    // "test": TestCmd,
    // "cover": CoverCmd,
    // "start": StartCmd,
    // "versions": VersionsCmd,
    // "deploy": DeployCmd,
}

var Aliases = []struct {
    alias string
    dest string
}{
    {"avweb", "/av_umbrella/apps/av_web"},
    {"avmain", "/av_umbrella/apps/av"},
    {"avassets", "/assets"},
    {"avnodeapp", "/nodeapp"},
}
