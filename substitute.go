package main

import (
	"fmt"
	"os/exec"
)

type Substitute struct {
	searchPattern  string
	replacePattern string
	paths          []string
}

func (s *Substitute) Run() ([]byte, error) {
	return s.command().CombinedOutput()
}

func (s *Substitute) grep() *exec.Cmd {
	grepArgs := []string{"grep", "--extended-regexp", "--files-with-matches", s.searchPattern}
	if len(s.paths) > 0 {
		grepArgs = append(grepArgs, s.paths...)
	}
	return exec.Command("git", grepArgs...)
}

func (s *Substitute) sed() *exec.Cmd {
	search := fmt.Sprintf("s/%s/%s/g", s.searchPattern, s.replacePattern)
	return exec.Command("xargs", "sed", "--regexp-extended", "--in-place", search)
}

func (s *Substitute) command() *exec.Cmd {
	grep := s.grep()
	sed := s.sed()
	grepOut, _ := grep.StdoutPipe()
	grep.Start()
	sed.Stdin = grepOut
	return sed
}
