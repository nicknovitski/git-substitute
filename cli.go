package main

import (
	"fmt"
	"os/exec"
	"os"
)

type CLI struct {
	searchPattern string
	replacePattern string
	paths []string
}

func cliFromDocopts(arguments map[string]interface{}) *CLI {
	return &CLI{
		searchPattern: arguments["<search-pattern>"].(string),
		replacePattern: arguments["<replace-pattern>"].(string),
		paths: arguments["<path>"].([]string),
	}
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
	grepArgs := []string{"grep", "-El", fmt.Sprintf("%s", cli.searchPattern)}
	if len(cli.paths) > 0 {
	  grepArgs = append(grepArgs, cli.paths...)
	}
	return exec.Command("git", grepArgs...)
}

func (cli *CLI) sed() *exec.Cmd {
	search := fmt.Sprintf("s/%s/%s/g", cli.searchPattern, cli.replacePattern)
	return exec.Command("xargs", "sed", "-ri", search)
}
