package main

import (
	"regexp"
)


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

func regex(pattern string, syntax regexSyntax) *regexp.Regexp {
	if syntax == perl {
		return regexp.MustCompile(regularizedPattern(pattern, syntax))
	} else {
		return regexp.MustCompilePOSIX(regularizedPattern(pattern, syntax))
	}
}

func regularizedPattern(pattern string, syntax regexSyntax) string {
	if syntax == basic {
		return escapeMetacharacters(pattern)
	} else {
		return pattern
	}
}
