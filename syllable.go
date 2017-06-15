package main

import (
	"regexp"
	"strings"
)

var r, _ = regexp.Compile("([a-zA-Z]+)")

type Syllable struct {
	text     string
	isPrefix bool
	isSuffix bool
}

func SyllableFactory(raw string) Syllable {
	raw = strings.TrimSpace(raw)
	return Syllable{text: stripMetadata(raw),
		isPrefix: strings.HasPrefix(raw, "-"),
		isSuffix: strings.HasPrefix(raw, "+")}
}

func stripMetadata(raw string) string {
	return r.FindStringSubmatch(raw)[0]
}
