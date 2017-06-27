package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Goname", func() {
	Describe("Given a request for a name with 3 syllables", func() {
		Context("With a request for a name with 3 syllables", func() {
			It("should return something", func() {
				name := New(RomanMap)
				Expect(name.Compose(4)).To(Equal("notrandom"))
			})
		})
	})
})
