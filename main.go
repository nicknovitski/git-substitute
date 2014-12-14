package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"os"
)

func main() {
	usage := `Usage:
  git-substitute <search-pattern> <replace-pattern> [<path>...]
  git-substitute -h | --help
  git-substitute -V | --version

Options:
  -h --help        Show this screen.
  -V --version     Show version.`

	arguments, _ := docopt.Parse(usage, nil, true, "git-substitute 1.1.1", false)
	command := &substitute{
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
