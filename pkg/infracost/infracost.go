package infracost

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/bradmccoydev/tfval/model"
)

func CheckIfSpendIsWithinBudget(infracostReport []byte, monthlyBudget float64) string {
	var breakdown model.InfracostBreakdown
	json.Unmarshal(infracostReport, &breakdown)

	monthlyCost, err := strconv.ParseFloat(breakdown.TotalMonthlyCost, 6)
	if err != nil {
		fmt.Println(err)
	}

	infracostPass := false
	if monthlyBudget > monthlyCost {
		infracostPass = true
	}

	return fmt.Sprintf("{\"infracost_pass\":%t,\"infracost_monthly_budget\":%f,\"infracost_total_monthly_cost\":%f", infracostPass, monthlyBudget, monthlyCost)

}
