package goname_test

import (
	. "github.com/folkengine/goname"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Goname", func() {
	Describe("Given a request for a name with 3 syllables", func() {
		Context("With a request for a name with 3 syllables", func() {
			It("should return something", func() {
				Expect(Generate(4)).To(Equal("notrandom"))
			})
		})
	})
})
