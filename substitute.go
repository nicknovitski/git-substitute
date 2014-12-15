package main

import (
	"fmt"
	"os/exec"
)

type substitute struct {
	searchPattern  string
	replacePattern string
	paths          []string
}

func (s *substitute) Run() ([]byte, error) {
	return s.command().CombinedOutput()
}

func (s *substitute) grep() *exec.Cmd {
	grepArgs := []string{"grep", "--extended-regexp", "--files-with-matches", s.searchPattern}
	if len(s.paths) > 0 {
		grepArgs = append(grepArgs, s.paths...)
	}
	return exec.Command("git", grepArgs...)
}

func (s *substitute) sed() *exec.Cmd {
	return exec.Command("xargs", "sed", "-E", "--in-place", s.sedSubCommand())
}

func (s *substitute) sedSubCommand() string {
	return fmt.Sprintf("s/%s/%s/g", s.searchPattern, s.replacePattern)
}

func (s *substitute) command() *exec.Cmd {
	grep := s.grep()
	sed := s.sed()
	grepOut, _ := grep.StdoutPipe()
	grep.Start()
	sed.Stdin = grepOut
	return sed
}
