package main

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var r, _ = regexp.Compile("(\\p{L}+)")
var consonants = "bʙβcdɗɖðfghjklłmɱnɳpqrstvwxyz"
var vowels = "iyɨʉɯuɪʏʊɯʊeøɘɵɤoøəɵɤoɛœɜɞʌɔæɐɞaɶäɒɑ"

// Syllable is an atomic building block used to construct names.
type Syllable struct {
	text                      string
	isPrefix                  bool
	isSuffix                  bool
	nextStartsWithConsonant   bool
	nextStartsWithVowel       bool
	previousEndsWithConsonant bool
	previousEndsWithVowel     bool
}

func (syllable Syllable) consonantFirst() bool {
	first, _ := utf8.DecodeRuneInString(syllable.text)
	return strings.Contains(consonants, strings.ToLower(string(first)))
}

func (syllable Syllable) vowelFirst() bool {
	first, _ := utf8.DecodeRuneInString(syllable.text)
	return strings.Contains(vowels, strings.ToLower(string(first)))
}

// SyllableFactory is a static factory method, generating Syllable object.
func SyllableFactory(raw string) Syllable {
	raw = strings.TrimSpace(raw)
	return Syllable{text: stripMetadata(raw),
		isPrefix:                  strings.HasPrefix(raw, "-"),
		isSuffix:                  strings.HasPrefix(raw, "+"),
		nextStartsWithConsonant:   strings.Contains(raw, "+c"),
		nextStartsWithVowel:       strings.Contains(raw, "+v"),
		previousEndsWithConsonant: strings.Contains(raw, "-c"),
		previousEndsWithVowel:     strings.Contains(raw, "-v")}
}

func stripMetadata(raw string) string {
	return r.FindStringSubmatch(raw)[0]
}
