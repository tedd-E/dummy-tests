package dummy_tests_test

import (
	"context"
	"fmt"
	"github.com/onsi/ginkgo/v2"
	"sync"
)

var (
	// Mutex for thread safety while appending results
	mu       sync.Mutex
	results  []string // Holds test results for this node
	nodeData []string // Consolidated results across nodes
)

var _ = ginkgo.SynchronizedAfterSuite(
	// Aggregation function (runs only on the primary node)
	func(ctx context.Context) {
		test := ctx.(ginkgo.SpecContext)
		report := test.SpecReport()
		//report := ginkgo.CurrentSpecReport()
		fmt.Printf("2Suite Report: %+v\n", report)
	},
	// Node-specific function (runs on all nodes)
	func(ctx context.Context) {
		fmt.Printf("test\n")
		//report := ginkgo.CurrentSpecReport()
		test := ctx.(ginkgo.SpecContext)
		report := test.SpecReport()

		fmt.Printf("Suite Report: %+v\n", report)

	},
)
