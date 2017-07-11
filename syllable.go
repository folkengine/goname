package goname

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
	raw                       string
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

func (syllable Syllable) consonantLast() bool {
	first, _ := utf8.DecodeLastRuneInString(syllable.text)
	return strings.Contains(consonants, strings.ToLower(string(first)))
}

func (syllable Syllable) vowelLast() bool {
	first, _ := utf8.DecodeLastRuneInString(syllable.text)
	return strings.Contains(vowels, strings.ToLower(string(first)))
}

func (syllable Syllable) nextIncompatible(next Syllable) bool {
	vnc := syllable.nextStartsWithVowel && next.consonantFirst()
	cnv := syllable.nextStartsWithConsonant && next.vowelFirst()
	return vnc || cnv
}

func (syllable Syllable) previousIncompatible(next Syllable) bool {
	vlc := syllable.vowelLast() && next.previousEndsWithConsonant
	clv := syllable.consonantLast() && next.previousEndsWithVowel
	return vlc || clv
}

func (syllable Syllable) incompatible(next Syllable) bool {
	return syllable.nextIncompatible(next) || syllable.previousIncompatible(next)
}

func (syllable Syllable) compatible(next Syllable) bool {
	return !syllable.incompatible(next)
}

// SyllableFactory is a static factory method, generating Syllable object.
func SyllableFactory(raw string) Syllable {
	raw = strings.TrimSpace(raw)
	return Syllable{
		raw:                       raw,
		text:                      stripMetadata(raw),
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
