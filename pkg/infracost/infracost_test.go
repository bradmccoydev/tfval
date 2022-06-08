package infracost

import (
	"fmt"
	"testing"

	"github.com/bradmccoydev/tfval/pkg/utils"
)

func TestGetTfScore(t *testing.T) {
	testCases := []struct {
		monthlyBudget           float64
		infracostReportLocation string
		err                     error
	}{
		{2000.00, "mock.json", nil},
		{0.638240, "mock.json", nil},
	}
	for _, tc := range testCases {
		infracostReport := utils.ReadFile(tc.infracostReportLocation)
		infracostResponse := CheckIfSpendIsWithinBudget([]byte(infracostReport), tc.monthlyBudget)
		fmt.Println(infracostResponse)
	}
}
