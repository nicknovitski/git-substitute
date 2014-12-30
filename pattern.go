package main

import (
	"regexp"
)

type searchPattern struct {
	pattern string
	syntax  regexSyntax
}

func (s *searchPattern) string() string {
	return s.pattern
}

func (s *searchPattern) grepArg() string {
	if s.syntax == perl {
		return "--perl-regexp"
	} else if s.syntax == basic {
		return "--basic-regexp"
	} else {
		return "--extended-regexp"
	}
}

func (s *searchPattern) regexp() *regexp.Regexp {
	if s.syntax == perl {
		return regexp.MustCompile(s.goStyle())
	} else {
		return regexp.MustCompilePOSIX(s.goStyle())
	}
}

func (s *searchPattern) goStyle() string {
	if s.syntax == basic {
		return escapeMetacharacters(s.pattern)
	} else {
		return s.pattern
	}
}

func escapeMetacharacters(target string) string {
	metas := []string{`\?`, `\+`, `\|`}
	for _, meta := range metas {
		target = regexp.MustCompile(meta).ReplaceAllString(target, meta)
	}
	result := parensAndBrackets().ReplaceAllFunc([]byte(target), reverseEscape(`(`, `)`, `{`, `}`))
	return string(result)
}

func parensAndBrackets() *regexp.Regexp {
	return regexp.MustCompile(`\\?[\(\)\{\}]`)
}

func reverseEscape(matches ...string) func([]byte) []byte {
	escapedMatches := make([]string, len(matches))
	for i, match := range matches {
		escapedMatches[i] = `\` + match
	}

	return func(match []byte) []byte {
		for i, m := range matches {
			if m == string(match) {
				return []byte(escapedMatches[i])
			}
		}
		for i, m := range escapedMatches {
			if m == string(match) {
				return []byte(matches[i])
			}
		}
		return match
	}
}
