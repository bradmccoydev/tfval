package tfsec

import (
	"fmt"
	"testing"
)

func TestProduceVulnerabilityReport(t *testing.T) {
	report := ProduceVulnerabilityReport("mock.json")

	fmt.Println(report)
}
