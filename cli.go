package main

import (
	"fmt"
	"os/exec"
	"os"
)

type CLI struct {
	arguments map[string]interface{}
	paths []string
}

func newCLI(arguments map[string]interface{}) *CLI {
	return &CLI{arguments: arguments, paths: arguments["<path>"].([]string)}
}

func (cli *CLI) Run() {
	grep := cli.grep()
	sed := cli.sed()
	grepOut, _ := grep.StdoutPipe()
	grep.Start()
	sed.Stdin = grepOut
	output, err := sed.CombinedOutput()
	fmt.Print(string(output))
	if err != nil {
		os.Exit(1)
	}
}

func (cli *CLI) grep() *exec.Cmd {
	grepArgs := []string{"grep", "-El", fmt.Sprintf("%s", cli.arguments["<search-pattern>"])}
	if len(cli.paths) > 0 {
	  grepArgs = append(grepArgs, cli.paths...)
	}
	return exec.Command("git", grepArgs...)
}

func (cli *CLI) sed() *exec.Cmd {
	search := fmt.Sprintf("s/%s/%s/g", cli.arguments["<search-pattern>"], cli.arguments["<replace-pattern>"])
	return exec.Command("xargs", "sed", "-ri", search)
}