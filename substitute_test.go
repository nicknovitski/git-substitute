package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

func fileContains (filename string, contents string) bool{
	b, _ := ioutil.ReadFile(filename)
	return string(b) == contents
}

func TestReplaceInPath(t *testing.T) {
	c := &Substitute{
		searchPattern:  "foo",
		replacePattern: "bar",
		paths:          []string{"test1"},
	}
	ioutil.WriteFile("test1", []byte("foo"), 0644)
	defer os.Remove("test1")
	ioutil.WriteFile("test2", []byte("foo"), 0644)
	defer os.Remove("test2")

	exec.Command("git", "add", "test1", "test2").Run()
	defer exec.Command("git", "reset", "HEAD", "test1", "test2").Run()
	
	c.Run()
	if fileContains("test1", "foo") {
		t.Error("pattern not replaced in target")
	}
	if fileContains("test2", "bar") {
		t.Error("pattern replaced in non-target")
	}
}

func TestIgnoresFilesNotInGit(t *testing.T) {
	t.Skip("god i'm tired")
}

func TestFancyPatterns(t *testing.T) {
	t.Skip("I need to think of a good example first.")
}
