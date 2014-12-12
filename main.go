package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"os"
)

func main() {
	usage := `Git Substitute.

Usage:
  git-substitute <search-pattern> <replace-pattern> [<path> ...]
  git-substitute -h | --help
  git-substitute -V | --version

Options:
  -h --help     Show this screen.
  --version     Show version.`

	arguments, _ := docopt.Parse(usage, nil, true, "Git Substitute 1.0", false)
	command := &Substitute{
		searchPattern:  arguments["<search-pattern>"].(string),
		replacePattern: arguments["<replace-pattern>"].(string),
		paths:          arguments["<path>"].([]string),
	}
	out, err := command.Run()
	if err != nil {
		fmt.Print(string(out))
		os.Exit(1)
	}
}
