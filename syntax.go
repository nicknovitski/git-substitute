package main

type regexSyntax int

const (
	basic regexSyntax = iota
	extended
	perl
	fixed
)

func syntax(options map[string]interface{}) regexSyntax {
	if options["--basic-regexp"].(bool) {
		return basic
	} else if options["--perl-regexp"].(bool) {
		return perl
	} else if options["--fixed-strings"].(bool) {
		return fixed
	} else {
		return extended
	}
}
