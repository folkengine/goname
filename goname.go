package main

import (
	"bytes"
	"math/rand"
	"strings"
	"time"
)

var numberSyllablesPik = [...]int{2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 5}

func init() {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
}

// Goname name generating thing.
type Goname struct {
	preSyllables []Syllable
	midSyllables []Syllable
	surSyllables []Syllable
}

func (goname Goname) preSyllableSample() Syllable {
	return sampleSyllable(goname.preSyllables)
}

func (goname Goname) midSyllableSample() Syllable {
	return sampleSyllable(goname.midSyllables)
}

func (goname Goname) surSyllableSample() Syllable {
	return sampleSyllable(goname.surSyllables)
}

// NameSyllables returns a slice of Syllables based upon the number passed in.
func (goname Goname) NameSyllables(number int) []Syllable {
	var nameSyllables []Syllable
	preSyllable := goname.preSyllableSample()
	nameSyllables = append(nameSyllables, preSyllable)
	for i := 1; i < number-1; i++ {
		nameSyllables = append(nameSyllables, goname.midSyllableSample())
	}
	if number > 1 {
		nameSyllables = append(nameSyllables, goname.surSyllableSample())
	}
	return nameSyllables
}

// NameFromSyllables returns a Name from syllables.
func (goname Goname) NameFromSyllables(syllables []Syllable) string {
	var buffer bytes.Buffer
	for _, syllable := range syllables {
		buffer.WriteString(syllable.text)
	}
	return strings.Title(buffer.String())
}

// Name random name.
func (goname Goname) Name(number int) string {
	return goname.NameFromSyllables(goname.NameSyllables(number))
}

// NameRnd generates random syllable length name
func (goname Goname) NameRnd() string {
	return goname.Name(pickSyllablesCount())
}

// FirstLast generates a string with a first and last name from random syllables.
func (goname Goname) FirstLast() string {
	return goname.NameRnd() + " " + goname.NameRnd()
}

// New is a static factory method generating Gonames
func New(dialectMap map[string]Syllable) Goname {
	var first []Syllable
	var mid []Syllable
	var last []Syllable

	for _, v := range dialectMap {
		if v.isPrefix {
			first = append(first, v)
		} else if v.isSuffix {
			last = append(last, v)
		} else {
			mid = append(mid, v)
		}
	}
	return Goname{
		preSyllables: first,
		midSyllables: mid,
		surSyllables: last,
	}
}

func sampleSyllable(syllables []Syllable) Syllable {
	return syllables[rand.Intn(len(syllables))]
}

func pickSyllablesCount() int {
	return numberSyllablesPik[rand.Intn(len(numberSyllablesPik))]
}
