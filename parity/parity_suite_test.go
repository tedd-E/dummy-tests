package parity_test

import (
	fern "github.com/guidewire/fern-ginkgo-client/pkg/client"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tedd-E/dummy-tests/parity"
)

func TestEvenOdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EvenOdd Suite")
}

var _ = Describe("IsEven", Label("parity"), func() {
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

var _ = ReportAfterSuite("", func(report Report) {
	f := fern.New("Parity Tests",
		fern.WithBaseURL("http://localhost:8080/"),
	)

	err := f.Report("Parity Tests", report)

	Expect(err).To(BeNil(), "Unable to create reporter file")
})
