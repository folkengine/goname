package goname

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoname(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goname Suite")
}
