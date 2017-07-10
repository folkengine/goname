package main

import (
	_ "fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var sut = New(ElvenMap)

var _ = Describe("Goname", func() {
	Describe("When creating a new Goname entity", func() {
		Context("When passing in a particular language map", func() {
			It("All of the syllable fields should have the correct Syllables", func() {
				Expect(sut.preSyllables[0].isPrefix).To(Equal(true))
				Expect(sut.midSyllables[0].isPrefix).To(Equal(false))
				Expect(sut.midSyllables[0].isSuffix).To(Equal(false))
				Expect(sut.surSyllables[0].isPrefix).To(Equal(false))
				Expect(sut.surSyllables[0].isSuffix).To(Equal(true))
			})
		})

		Context("When generating a word with NameSyllables", func() {
			It("Should return a pre syllable as the first syllable", func() {
				name := sut.NameSyllables(1)
				Expect(languagesContain(name[0], sut.preSyllables)).To(Equal(true))
				Expect(len(name)).To(Equal(1))
			})
			It("Should return a sur syllable as the last syllable", func() {
				name := sut.NameSyllables(2)
				Expect(languagesContain(name[1], sut.surSyllables)).To(Equal(true))
			})
			It("Should return a pre, a mid, and a sur syllable if there are three", func() {
				name := sut.NameSyllables(3)
				Expect(languagesContain(name[0], sut.preSyllables)).To(Equal(true))
				Expect(languagesContain(name[1], sut.midSyllables)).To(Equal(true))
				Expect(languagesContain(name[2], sut.surSyllables)).To(Equal(true))
			})
			It("Should return a pre, a mid, and a sur syllable if there are three", func() {
				name := sut.NameSyllables(4)
				Expect(languagesContain(name[0], sut.preSyllables)).To(Equal(true))
				Expect(languagesContain(name[1], sut.midSyllables)).To(Equal(true))
				Expect(languagesContain(name[2], sut.midSyllables)).To(Equal(true))
				Expect(languagesContain(name[3], sut.surSyllables)).To(Equal(true))
			})
		})

	})
})

func languagesContain(syllable Syllable, syllables []Syllable) bool {
	for _, s := range syllables {
		if syllable == s {
			return true
		}
	}
	return false
}
