package main

import (
	"regexp"
)

var wordRe = regexp.MustCompile(`(['\w]+(?:-['\w]+)*)`)

func Parse(input string) []string {
	captures := wordRe.FindAllStringSubmatch(input, -1)
	if captures == nil {
		return nil
	}
	var results = make([]string, 0, len(captures))
	for _, capture := range captures {
		s := capture[1]
		if s == `'` {
			continue
		}
		results = append(results, s)
	}
	return results
}
