package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)


var numberSyllablesPik = [...]int {2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 5}

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

// New is a static factory method generating NameGenerators
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

// Compose random name.
func (goname Goname) Compose(number int) string {
	preSyllable := goname.preSyllableSample()
	if number < 2 {
		return strings.Title(preSyllable.text)
	}
	return "notrandom"
}

func (goname Goname) ComposeRnd() string {
	return goname.Compose(pickSyllablesCount())
}

func sampleSyllable(syllables []Syllable) Syllable {
	return syllables[rand.Intn(len(syllables))]
}

func pickSyllablesCount() int {
	return numberSyllablesPik[rand.Intn(len(numberSyllablesPik))]
}

func main() {
	goname := New(ElvenMap)
	fmt.Println(goname.ComposeRnd())
}
