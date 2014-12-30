package main

import (
	"regexp"
	"strconv"
)

type replacePattern struct {
	pattern string
}

func (r *replacePattern) highestBackreferenceNumber() int {
	backrefs := regexp.MustCompile(`\\\d`).FindAllString(r.pattern, -1)
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

func (r *replacePattern) goStyle() []byte {
	return []byte(regexp.MustCompile(`\\(\d)`).ReplaceAllString(r.pattern, `$$$1`))
}

func (r *replacePattern) string() string {
	return r.pattern
}
