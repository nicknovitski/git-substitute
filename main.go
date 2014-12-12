package main

import (
	"github.com/docopt/docopt-go"
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
	newCLI(arguments).Run()
}
