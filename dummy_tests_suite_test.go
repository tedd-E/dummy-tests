package dummy_tests_test

import (
	fern "github.com/guidewire-oss/fern-ginkgo-client/pkg/client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	_ "github.com/tedd-E/dummy-tests/addition"
	_ "github.com/tedd-E/dummy-tests/parity"
	_ "github.com/tedd-E/dummy-tests/strings"
	"testing"
)

func TestDummyTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DummyTests Suite")
}

var _ = ReportAfterSuite("", func(report Report) {
	f := fern.New("Dummy Tests",
		fern.WithBaseURL("http://localhost:8080/"),
	)

	err := f.Report("Dummy Tests", report)

	Expect(err).To(BeNil(), "Unable to create reporter file")
})
