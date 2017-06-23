package main

import "fmt"

type NameGenerator struct {
	pre_syllables 	[]Syllable
	mid_syllables	[]Syllable
	sur_syllables  	[]Syllable
}

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
		pre_syllables: first,
		mid_syllables: mid,
		sur_syllables: last}
}

//Generate random name.
func Generate(number int) string {
	return "notrandom"
}

func main() {
	fmt.Println(Generate(5))
}
