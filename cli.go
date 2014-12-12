package main

import (
	"fmt"
	"os/exec"
)

type CLI struct {
	arguments map[string]interface{}
}

func newCLI(arguments map[string]interface{}) *CLI {
	return &CLI{arguments}
}

func (cli *CLI) Run() {
	grep := cli.grep()
	sed := cli.sed()
	grepOut, _ := grep.StdoutPipe()
	grep.Start()
	sed.Stdin = grepOut
	output, _ := sed.CombinedOutput()
	fmt.Print(string(output))
}

func (cli *CLI) grep() *exec.Cmd {
	grepArgs := []string{"grep", "-El", fmt.Sprintf("%s", cli.arguments["<search-pattern>"])}
	if cli.arguments["<path>"] != nil {
		grepArgs = append(grepArgs, cli.arguments["<path>"].(string))
	}
	return exec.Command("git", grepArgs...)
}

func (cli *CLI) sed() *exec.Cmd {
	search := fmt.Sprintf("s/%s/%s/g", cli.arguments["<search-pattern>"], cli.arguments["<replace-pattern>"])
	return exec.Command("xargs", "sed", "-ri", search)
}
