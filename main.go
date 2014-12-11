package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"os/exec"
)

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
	grepArgs := []string{"grep", "-El", fmt.Sprintf("%s", arguments["<search-pattern>"])}
	if arguments["<path>"] != nil {
		grepArgs = append(grepArgs, arguments["<path>"].(string))
	}
	grep := exec.Command("git", grepArgs...)
	sed := exec.Command("xargs", "sed", "-ri", fmt.Sprintf("s/%s/%s/g", arguments["<search-pattern>"], arguments["<replace-pattern>"]))
	grepOut, _ := grep.StdoutPipe()
	grep.Start()
	sed.Stdin = grepOut
	output, _ := sed.CombinedOutput()
	fmt.Print(string(output))
}
