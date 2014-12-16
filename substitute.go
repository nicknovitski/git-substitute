package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type substitute struct {
	searchPattern  string
	replacePattern string
	paths          []string
}

func (s *substitute) Run() ([]byte, error) {
	return s.command().CombinedOutput()
}

func (s *substitute) matchingFiles() []string {
	grepArgs := []string{"grep", "--extended-regexp", "--files-with-matches", s.searchPattern}
	if len(s.paths) > 0 {
		grepArgs = append(grepArgs, s.paths...)
	}
	output, err := exec.Command("git", grepArgs...).CombinedOutput()
	if err != nil {
		fmt.Println(output)
		os.Exit(1)
	}
	splitOut := strings.Split(string(output), "\n")
	return splitOut[:len(splitOut)-1]
}

func (s *substitute) sed(files []string) *exec.Cmd {
	sedArgs := []string{"-E", "-i", fmt.Sprintf("-e %s", s.sedSubCommand())}
	if len(files) > 0 {
		sedArgs = append(sedArgs, files...)
	}
	return exec.Command("sed", sedArgs...)
}

func (s *substitute) sedSubCommand() string {
	return fmt.Sprintf("s/%s/%s/g", s.searchPattern, s.replacePattern)
}

func (s *substitute) command() *exec.Cmd {
	return s.sed(s.matchingFiles())
}
