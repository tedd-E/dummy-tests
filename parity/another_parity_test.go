package parity_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tedd-E/dummy-tests/parity"
)

var _ = Describe("IsEvenNegative", Label("parity", "negative"), func() {
	It("should return true for a negative even number", func() {
		Expect(parity.IsEven(-4)).To(BeTrue())
	})

	It("should return false for a negative odd number", func() {
		Expect(parity.IsEven(-3)).To(BeFalse())
	})
})
