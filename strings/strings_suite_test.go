package strings_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tedd-E/dummy-tests/strings"
	"testing"
)

func TestStrings(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Strings Suite")
}

var _ = Describe("StringLength", func() {
	It("should return 0 for an empty string", func() {
		Expect(strings.StringLength("")).To(Equal(0))
	})

	It("should return the correct length for a non-empty string", func() {
		Expect(strings.StringLength("hello")).To(Equal(5))
	})

	It("should return the correct length for a string with spaces", func() {
		Expect(strings.StringLength("hello world")).To(Equal(11))
	})
})
