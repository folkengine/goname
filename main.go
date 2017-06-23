package main

import "fmt"

// NameGenerator namge generating thing.
type NameGenerator struct {
	preSyllables []Syllable
	midSyllables []Syllable
	surSyllables []Syllable
}

// Goname is a static factory method generating NameGenerators
func Goname(dialectMap map[string]Syllable) NameGenerator {
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
	return NameGenerator{
		preSyllables: first,
		midSyllables: mid,
		surSyllables: last,
	}
}

//Generate random name.
func Generate(number int) string {
	return "notrandom"
}

func main() {
	fmt.Println(Generate(5))
}
