package main

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Syllable", func() {

	Describe("For raw strings", func() {
		Context("For all of the types", func() {
			It("is should separate out the actual string", func() {
				Expect(SyllableFactory("-an").text).To(Equal("an"))
				Expect(SyllableFactory("-ang +v").text).To(Equal("ang"))
				Expect(SyllableFactory("  -ansr +v").text).To(Equal("ansr"))
				Expect(SyllableFactory("-cael   ").text).To(Equal("cael"))
				Expect(SyllableFactory("-dae +c").text).To(Equal("dae"))
			})
		})
	})

	Describe("For prefix syllables", func() {
		Context("With syllables that are prefixes", func() {
			It("should return true", func() {
				syllables := [5]string{"-an", "-ang +v", "  -ansr +v", "-cael   ", "-dae +c"}
				for _, raw := range syllables {
					syllable := SyllableFactory(raw)
					Expect(syllable.isPrefix).To(Equal(true),
						fmt.Sprintf("SyllableFactory(%s).isPrefix %t", syllable.text, syllable.isPrefix))
				}
			})
		})
		Context("With syllables that are not prefixes", func() {
			It("should return false", func() {
				syllables := [8]string{"que -v +c", "ria", "rail", "ther", "thus", "thi", "san", "+ael -c"}
				for _, raw := range syllables {
					syllable := SyllableFactory(raw)
					Expect(syllable.isPrefix).To(Equal(false),
						fmt.Sprintf("SyllableFactory(%s).isPrefix %t", raw, syllable.isPrefix))
				}
			})
		})
	})

	Describe("For suffix syllables", func() {
		Context("With syllables that are suffixes", func() {
			It("should return true", func() {
				syllables := [7]string{"+ath", "+ess", "+san", "+yth", "+las", "+lian", "+evar"}
				for _, raw := range syllables {
					syllable := SyllableFactory(raw)
					Expect(syllable.isSuffix).To(Equal(true),
						fmt.Sprintf("SyllableFactory(%s).isSuffix %t", syllable.text, syllable.isSuffix))
				}
			})
		})
		Context("With syllables that are not suffixes", func() {
			It("should return false", func() {
				syllables := [8]string{"-que -v +c", "ria", "rail", "ther", "thus", "thi", "san", "ael -c"}
				for _, raw := range syllables {
					syllable := SyllableFactory(raw)
					Expect(syllable.isSuffix).To(Equal(false),
						fmt.Sprintf("SyllableFactory(%s).isSuffix %t", raw, syllable.isSuffix))
				}
			})
		})
	})

	Describe("Next syllable", func() {
		Context("must start with a consonant", func() {
			It("should return true if marked", func() {
				syllables := [2]string{"-ang +c", "-ansr -v +c"}
				for _, raw := range syllables {
					syllable := SyllableFactory(raw)
					Expect(syllable.nextStartsWithConsonant).To(Equal(true),
						fmt.Sprintf("SyllableFactory(%s).nextStartsWithConsonant %t", syllable.text, syllable.nextStartsWithConsonant))
				}
			})
			It("should return false if not marked", func() {
				syllables := [3]string{"-ang -v", "-ansr -v -c", "-yada"}
				for _, raw := range syllables {
					syllable := SyllableFactory(raw)
					Expect(syllable.nextStartsWithConsonant).To(Equal(false),
						fmt.Sprintf("SyllableFactory(%s).nextStartsWithConsonant %t", raw, syllable.nextStartsWithConsonant))
				}
			})
		})

		Context("next syllable must start with vowel", func() {
			It("should return true if marked", func() {
				syllables := [2]string{"-ang +v", "-ansr +v -c"}
				for _, raw := range syllables {
					syllable := SyllableFactory(raw)
					Expect(syllable.nextStartsWithVowel).To(Equal(true),
						fmt.Sprintf("SyllableFactory(%s).nextStartsWithVowel %t", syllable.text, syllable.nextStartsWithVowel))
				}
			})
			It("should return false if not marked", func() {
				syllables := [3]string{"-ang -v", "-ansr -v +c", "-yada"}
				for _, raw := range syllables {
					syllable := SyllableFactory(raw)
					Expect(syllable.nextStartsWithVowel).To(Equal(false),
						fmt.Sprintf("SyllableFactory(%s).nextStartsWithVowel %t", raw, syllable.nextStartsWithVowel))
				}
			})
		})
	})

	Describe("Previous syllable", func() {
		Context("must end with a consonant", func() {
			It("should return true if marked", func() {
				syllables := [2]string{"-ang -c", "-ansr +v -c"}
				for _, raw := range syllables {
					syllable := SyllableFactory(raw)
					Expect(syllable.previousEndsWithConsonant).To(Equal(true),
						fmt.Sprintf("SyllableFactory(%s).previousEndsWithConsonant %t", syllable.text, syllable.previousEndsWithConsonant))
				}
			})
			It("should return false if not marked", func() {
				syllables := [3]string{"-ang +v", "-ansr -v +c", "-yada"}
				for _, raw := range syllables {
					syllable := SyllableFactory(raw)
					Expect(syllable.previousEndsWithConsonant).To(Equal(false),
						fmt.Sprintf("SyllableFactory(%s).previousEndsWithConsonant %t", raw, syllable.previousEndsWithConsonant))
				}
			})
		})
	})
})
