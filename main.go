package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"os/exec"
)

func main() {
	usage := `Git Substitute.

Usage:
  git-substitute <search-pattern> <replace-pattern>
  git-substitute -h | --help
  git-substitute -V | --version

Options:
  -h --help     Show this screen.
  --version     Show version.`

	arguments, _ := docopt.Parse(usage, nil, true, "Git Substitute 1.0", false)
	grep := exec.Command("git", "grep", "-El", fmt.Sprintf("%s", arguments["<search-pattern>"]))
	sed := exec.Command("xargs", "sed", "-ri", fmt.Sprintf("s/%s/%s/g", arguments["<search-pattern>"], arguments["<replace-pattern>"]))
	grepOut, _ := grep.StdoutPipe()
	grep.Start()
	sed.Stdin = grepOut
	output, _ := sed.CombinedOutput()
	fmt.Print(string(output))
}
