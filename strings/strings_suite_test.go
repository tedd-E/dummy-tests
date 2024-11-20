package strings_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tedd-E/dummy-tests/strings"
	"testing"
)

func TestStrings(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Strings Suite", Label("string", "non-empty", "hi", "wough", "hello", "kubevela", "acceptance"))
}

var _ = Describe("StringLength", Label("abcde"), func() {
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

//var _ = ReportAfterSuite("", func(report Report) {
//	f := fern.New("String Tests",
//		fern.WithBaseURL("http://localhost:8080/"),
//	)
//
//	err := f.Report("String Tests", report)
//
//	Expect(err).To(BeNil(), "Unable to create reporter file")
//})
