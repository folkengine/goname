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
})
