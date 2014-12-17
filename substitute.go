package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type substitute struct {
	searchPattern  string
	replacePattern string
	paths          []string
}

func (s *substitute) Run() error {
	for _, file := range s.matchingFiles() {
		contents, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(file, s.replace(contents), 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *substitute) replace(source []byte) []byte {
	// Go uses $N for backreferences, not \N
	goPattern := regexp.MustCompile(`\\(\d)`).ReplaceAllString(s.replacePattern, `$$$1`)
	return s.regex().ReplaceAll(source, []byte(goPattern))
}

func (s *substitute) regex() *regexp.Regexp {
	return regexp.MustCompilePOSIX(s.searchPattern)
}

func (s *substitute) matchingFiles() []string {
	grepArgs := []string{"grep", "--extended-regexp", "--files-with-matches", s.searchPattern}
	if len(s.paths) > 0 {
		grepArgs = append(grepArgs, s.paths...)
	}
	output, err := exec.Command("git", grepArgs...).CombinedOutput()
	if err != nil {
		if len(output) != 0 {
			fmt.Println(output)
		}
		os.Exit(1)
	}
	splitOut := strings.Split(string(output), "\n")
	return splitOut[:len(splitOut)-1]
}
