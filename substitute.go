package main

import (
	"errors"
	"io/ioutil"
	"strings"
)

type substitute struct {
	searchPattern  string
	replacePattern *replacePattern
	paths          []string
	syntax         regexSyntax
}

func Sub(searchPattern string, repl string, paths []string, syntax regexSyntax) *substitute {
	return &substitute{
		searchPattern:  searchPattern,
		replacePattern: &replacePattern{pattern: repl},
		paths:          paths,
		syntax:         syntax,
	}
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
	if s.replacePattern.highestBackreferenceNumber() > regex(s.searchPattern, s.syntax).NumSubexp() {
		return errors.New("backreference without matching group expression")
	} else {
		return nil
	}
}

func (s *substitute) replace(source []byte) []byte {
	if s.syntax == fixed {
		return s.literalReplace(source)
	}
	return regex(s.searchPattern, s.syntax).ReplaceAll(source, s.replacePattern.goStyle())
}

func (s *substitute) literalReplace(source []byte) []byte {
	return []byte(strings.Replace(string(source), s.searchPattern, s.replacePattern.string(), -1))
}
