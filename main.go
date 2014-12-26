package main

import (
	"github.com/docopt/docopt-go"
	"os"
)

func syntax(options map[string]interface{}) regexSyntax {
	if options["--basic-regexp"].(bool) {
		return basic
	} else if options["--perl-regexp"].(bool) {
		return perl
	} else {
		return extended
	}
}
func main() {
	gitStatus()
	usage := `Usage:
  git-substitute [options] <search-pattern> <replace-pattern> [<path>...]
  git-substitute -h | --help
  git-substitute -V | --version

Options:
  -G --basic-regexp     Use basic POSIX regular expressions.
  -E --extended-regexp  Use extended POSIX regular expressions.
  -P --perl-regexp      Use Perl-compatible regular expressions.
  -h --help             Show this screen.
  -V --version          Show version.`

	arguments, _ := docopt.Parse(usage, nil, true, "git-substitute 1.4.0", false)
	command := &substitute{
		searchPattern:  arguments["<search-pattern>"].(string),
		replacePattern: arguments["<replace-pattern>"].(string),
		paths:          arguments["<path>"].([]string),
		syntax:         syntax(arguments),
	}
	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}
