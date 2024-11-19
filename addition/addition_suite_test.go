package addition_test

import (
	fern "github.com/guidewire-oss/fern-ginkgo-client/pkg/client"
	"github.com/tedd-E/dummy-tests/addition"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMath(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Math Suite", Label("sum", "addition", "hello", "hi", "wough", "atmos", "keda", "acceptance"))
}

var f *fern.FernApiClient

var _ = BeforeSuite(func() {
	f = fern.New("Example Test",
		fern.WithBaseURL("http://localhost:8080/"),
	)
	f.InitializeTestRun("Example Project") // Initializes the aggregated report
})

var _ = Describe("Sum", Label("sum"), func() {
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

//var _ = ReportAfterSuite("", func(report Report) {
//	f := fern.New("Addition Tests",
//		fern.WithBaseURL("http://localhost:8080/"),
//	)
//
//	err := f.Report("Addition Tests", report)
//
//	Expect(err).To(BeNil(), "Unable to create reporter file")
//})

var _ = ReportAfterSuite("", func(report Report) {
	err := f.Report("example test", report)
	Expect(err).To(BeNil(), "Unable to report suite results")
})

// After all tests are complete, submit the final aggregated report
var _ = AfterSuite(func() {
	err := f.SubmitFinalReport() // Sends the final aggregated report
	Expect(err).To(BeNil(), "Unable to submit final report")
})
