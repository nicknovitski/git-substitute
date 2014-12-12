package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"os/exec"
)

func grep(arguments map[string]interface{}) *exec.Cmd {
	grepArgs := []string{"grep", "-El", fmt.Sprintf("%s", arguments["<search-pattern>"])}
	if arguments["<path>"] != nil {
		grepArgs = append(grepArgs, arguments["<path>"].(string))
	}
	return exec.Command("git", grepArgs...)
}

func sed(arguments map[string]interface{}) *exec.Cmd {
    search := fmt.Sprintf("s/%s/%s/g", arguments["<search-pattern>"], arguments["<replace-pattern>"])
	return exec.Command("xargs", "sed", "-ri", search)
}

func main() {
	usage := `Git Substitute.

Usage:
  git-substitute <search-pattern> <replace-pattern> [<path>]
  git-substitute -h | --help
  git-substitute -V | --version

Options:
  -h --help     Show this screen.
  --version     Show version.`

	arguments, _ := docopt.Parse(usage, nil, true, "Git Substitute 1.0", false)
	grep := grep(arguments)
	sed := sed(arguments)
	grepOut, _ := grep.StdoutPipe()
	grep.Start()
	sed.Stdin = grepOut
	output, _ := sed.CombinedOutput()
	fmt.Print(string(output))
}
