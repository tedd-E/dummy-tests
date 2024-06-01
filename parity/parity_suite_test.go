package parity_test

import (
	"github.com/tedd-E/dummy-tests/parity"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEvenOdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EvenOdd Suite")
}

var _ = Describe("IsEven", func() {
	It("should return true for 0", func() {
		Expect(parity.IsEven(0)).To(BeTrue())
	})

	It("should return true for a positive even number", func() {
		Expect(parity.IsEven(4)).To(BeTrue())
	})

	It("should return false for a positive odd number", func() {
		Expect(parity.IsEven(3)).To(BeFalse())
	})

	It("should return true for a negative even number", func() {
		Expect(parity.IsEven(-2)).To(BeTrue())
	})

	It("should return false for a negative odd number", func() {
		Expect(parity.IsEven(-3)).To(BeFalse())
	})
})
