package main

import (
	"errors"
	"io/ioutil"
	"regexp"
	"strconv"
)

type regexSyntax int

const (
	basic regexSyntax = iota
	extended
	perl
)

type substitute struct {
	searchPattern  string
	replacePattern string
	paths          []string
	syntax         regexSyntax
}

func (s *substitute) Run() error {
	if bErr := s.backreferenceError(); bErr != nil {
		return bErr
	}
	for _, file := range filesMatching(s.searchPattern, s.syntax, s.paths) {
		contents, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		if err = ioutil.WriteFile(file, s.replace(contents), 0777); err != nil {
			return err
		}
	}
	return nil
}

func (s *substitute) backreferenceError() error {
	if highestNumberedBackreference(s.replacePattern) > s.regex().NumSubexp() {
		return errors.New("backreference without matching group expression")
	} else {
		return nil
	}
}

func highestNumberedBackreference(pattern string) int {
	backrefs := regexp.MustCompile(`\\\d`).FindAllString(pattern, -1)
	result := 0
	for _, backref := range backrefs {
		if curr := backreferenceNumber(backref); curr > result {
			result = curr
		}
	}
	return result
}

func backreferenceNumber(backref string) int {
	num, _ := strconv.Atoi(backref[1:])
	return num
}

func (s *substitute) replace(source []byte) []byte {
	return s.regex().ReplaceAll(source, []byte(s.goReplacePattern()))
}

func (s *substitute) goReplacePattern() string {
	// Go uses $N for backreferences, not \N
	return regexp.MustCompile(`\\(\d)`).ReplaceAllString(s.replacePattern, `$$$1`)
}

func (s *substitute) regex() *regexp.Regexp {
	if s.syntax == perl {
		return regexp.MustCompile(s.regularizedSearchPattern())
	} else {
		return regexp.MustCompilePOSIX(s.regularizedSearchPattern())
	}
}

func (s *substitute) regularizedSearchPattern() string {
	if s.syntax == basic {
		return escapeMetacharacters(s.searchPattern)
	} else {
		return s.searchPattern
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
