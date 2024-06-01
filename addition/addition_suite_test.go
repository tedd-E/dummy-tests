package addition_test

import (
	"github.com/tedd-E/dummy-tests/addition"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMath(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Math Suite")
}

var _ = Describe("Sum", func() {
	It("should return 0 when both inputs are 0", func() {
		Expect(addition.Sum(0, 0)).To(Equal(0))
	})

	It("should return the correct sum of positive numbers", func() {
		Expect(addition.Sum(2, 3)).To(Equal(5))
	})

	It("should return the correct sum when one input is negative", func() {
		Expect(addition.Sum(-2, 3)).To(Equal(1))
	})
})
