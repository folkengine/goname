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
		Context("must end with a vowel", func() {
			It("should return true if marked", func() {
				syllables := [2]string{"-ang -v", "-ansr -v +c"}
				for _, raw := range syllables {
					syllable := SyllableFactory(raw)
					Expect(syllable.previousEndsWithVowel).To(Equal(true),
						fmt.Sprintf("SyllableFactory(%s).previousEndsWithVowel %t", syllable.text, syllable.previousEndsWithVowel))
				}
			})
			It("should return false if not marked", func() {
				syllables := [3]string{"-ang +v", "-ansr +v -c", "-yada"}
				for _, raw := range syllables {
					syllable := SyllableFactory(raw)
					Expect(syllable.previousEndsWithVowel).To(Equal(false),
						fmt.Sprintf("SyllableFactory(%s).previousEndsWithVowel %t", raw, syllable.previousEndsWithVowel))
				}
			})
		})
	})

	Describe("First character", func() {
		Context("is a consonant", func() {
			It("should return true if it is", func() {
				syllables := [...]string{"-ng -c", "-sr +v -c", "-Yada", "łala", "βarg"}
				for _, raw := range syllables {
					syllable := SyllableFactory(raw)
					Expect(syllable.consonantFirst()).To(Equal(true),
						fmt.Sprintf("SyllableFactory(%s).consonantFirst() %t", syllable.text, syllable.consonantFirst()))
				}
			})
			It("should return false if it is not", func() {
				syllables := [...]string{"-ang +v", "-ansr -v +c", "Aadf","ʉnd", "œld", "øld"}
				for _, raw := range syllables {
					syllable := SyllableFactory(raw)
					Expect(syllable.consonantFirst()).To(Equal(false),
						fmt.Sprintf("SyllableFactory(%s).consonantFirst() %t", syllable.text, syllable.consonantFirst()))
				}
			})
		})

		Context("is a vowel", func() {
			It("should return true if it is", func() {
				syllables := [...]string{"-ang +v", "-ansr -v +c", "Aadf","ʉnd", "œld", "øld"}
				for _, raw := range syllables {
					syllable := SyllableFactory(raw)
					Expect(syllable.vowelFirst()).To(Equal(true),
						fmt.Sprintf("SyllableFactory(%s).vowelFirst() %t", syllable.text, syllable.vowelFirst()))
				}
			})
			It("should return false if it is not", func() {
				syllables := [...]string{"-ng -c", "-Sr +v -c", "-ðada", "łala", "βarg"}
				for _, raw := range syllables {
					syllable := SyllableFactory(raw)
					Expect(syllable.vowelFirst()).To(Equal(false),
						fmt.Sprintf("SyllableFactory(%s).vowelFirst() %t", syllable.text, syllable.vowelFirst()))
				}
			})
		})
	})

	Describe("Last character", func() {
		Context("is a consonant", func() {
			It("should return true if it is", func() {
				syllables := [...]string{"-ng -c", "-sr +v -c", "-ansr", "loβ"}
				for _, raw := range syllables {
					syllable := SyllableFactory(raw)
					Expect(syllable.consonantLast()).To(Equal(true),
						fmt.Sprintf("SyllableFactory(%s).consonantLast() %t", syllable.text, syllable.consonantLast()))
				}
			})
			It("should return false if it is not", func() {
				syllables := [...]string{"-yada", "ria", "thi", "thrɯ"}
				for _, raw := range syllables {
					syllable := SyllableFactory(raw)
					Expect(syllable.consonantLast()).To(Equal(false),
						fmt.Sprintf("SyllableFactory(%s).consonantLast() %t", syllable.text, syllable.consonantLast()))
				}
			})
		})

		Context("is a vowel", func() {
			It("should return true if it is", func() {
				syllables := [...]string{"-yada", "ria", "thi"}
				for _, raw := range syllables {
					syllable := SyllableFactory(raw)
					Expect(syllable.vowelLast()).To(Equal(true),
						fmt.Sprintf("SyllableFactory(%s).vowelLast() %t", syllable.text, syllable.vowelLast()))
				}
			})
			It("should return false if it is not", func() {
				syllables := [...]string{"-kan", "+emar", "+nes", "+nin", "dul", "ean -c", "el", "emar"}
				for _, raw := range syllables {
					syllable := SyllableFactory(raw)
					Expect(syllable.vowelLast()).To(Equal(false),
						fmt.Sprintf("SyllableFactory(%s).vowelLast() %t", syllable.text, syllable.vowelLast()))
				}
			})
		})
	})

	Describe("Incompatible", func() {
		Context("when determining if two syllables are incompatible", func() {
			It("it should return true when they are incompatible.", func() {
				Expect(ElvenMap["ae"]					.incompatible(SyllableFactory("lean -c"))).To(Equal(true))
				Expect(SyllableFactory("ria")		.incompatible(SyllableFactory("lean -c"))).To(Equal(true))
				Expect(SyllableFactory("at")		.incompatible(SyllableFactory("la -v"))).To(Equal(true))
				Expect(SyllableFactory("rail")		.incompatible(SyllableFactory("que -v +c"))).To(Equal(true))
				Expect(SyllableFactory("ae +c -c")	.incompatible(SyllableFactory("ael -c"))).To(Equal(true))
				Expect(SyllableFactory("ria")		.incompatible(SyllableFactory("ael -c"))).To(Equal(true))
			})
			It("it should return false when they are.", func() {
				Expect(SyllableFactory("lean -c")	.incompatible(SyllableFactory("ae +c -c"))).To(Equal(false))
				Expect(SyllableFactory("yada")		.incompatible(SyllableFactory("ria"))).To(Equal(false))
			})
		})
	})
	Describe("Compatible", func() {
		Context("when determining if two syllables are compatible", func() {
			It("it should return false when they are compatible.", func() {
				Expect(ElvenMap["ae"]					.compatible(SyllableFactory("lean -c"))).To(Equal(false))
				Expect(SyllableFactory("ria")		.compatible(SyllableFactory("lean -c"))).To(Equal(false))
				Expect(SyllableFactory("at")		.compatible(SyllableFactory("la -v"))).To(Equal(false))
				Expect(SyllableFactory("rail")		.compatible(SyllableFactory("que -v +c"))).To(Equal(false))
				Expect(SyllableFactory("ae +c -c")	.compatible(SyllableFactory("ael -c"))).To(Equal(false))
				Expect(SyllableFactory("ria")		.compatible(SyllableFactory("ael -c"))).To(Equal(false))
			})
			It("it should return true when they are.", func() {
				Expect(SyllableFactory("lean -c")	.compatible(SyllableFactory("ae +c -c"))).To(Equal(true))
				Expect(SyllableFactory("yada")		.compatible(SyllableFactory("ria"))).To(Equal(true))
			})
		})
	})
})
