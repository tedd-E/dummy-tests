package client

import fern "github.com/guidewire-oss/fern-ginkgo-client/pkg/client"

var FernClient = fern.New("Dummy Tests", fern.WithBaseURL("http://localhost:8080/"))
