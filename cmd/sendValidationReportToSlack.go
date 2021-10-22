package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/bradmccoydev/terraform-plan-validator/model"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var sendValidationReportToSlackCmd = &cobra.Command{
	Use:   "sendreport",
	Short: "Send Validation Report",
	Long:  `Send Terraform  validation Report to slack`,
	Run: func(cmd *cobra.Command, args []string) {
		err := sendValidationReportToSlack(args)
		if err != nil {
			fmt.Println(err)
		}
	},
}

var validationReportParams []string

func init() {
	rootCmd.AddCommand(sendValidationReportToSlackCmd)
	sendSlackMessageCmd.PersistentFlags().StringArrayVar(&reportParams, "payload", validationReportParams, "slackwebhook")
}

func sendValidationReportToSlack(args []string) error {
	prNumber := args[0]
	repoFullUrl := args[1]
	fileName := args[2]
	slackWebhook := args[3]

	report, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	var vulnerabilities model.Vulnerabilities
	json.Unmarshal([]byte(report), &vulnerabilities)

	if len(vulnerabilities.Results) > 0 {
		header := fmt.Sprintf(`{"blocks": [{"type": "header","text": {"type": "plain_text","text": ":cop: Pull Request %v Static Code Analysis Failed :cop:","emoji": true}}`, prNumber)
		footer := fmt.Sprintf(`,{"type": "divider"},{"type": "section","text": {"type": "mrkdwn","text": "View further details in the pull request:"},"accessory": {"type": "button","text": {"type": "plain_text","text": "View Pull Request","emoji": true},"value": "click_me_123","url": "%v","action_id": "button-action"}}`, repoFullUrl)
		body := fmt.Sprintf("%v%v", header, footer)

		for _, element := range vulnerabilities.Results {
			emoji := ":warning:"

			if element.Severity == "MEDIUM" {
				emoji = ":orange_book:"
			}

			if element.Severity == "CRITICAL" {
				emoji = ":red_envelope:"
			}

			impact := fmt.Sprintf(`,{"type": "section","text": {"type": "mrkdwn","text": "%v *%v*"}}`, emoji, element.Impact)
			resolution := fmt.Sprintf(`,{"type": "section","text": {"type": "mrkdwn","text": ":arrow_right: %v"}}`, element.Resolution)
			body = fmt.Sprintf(`%v%v%v`, body, impact, resolution)

		}

		body = fmt.Sprintf(`%v%v`, body, "]}")
		fmt.Println(body)

		method := "POST"

		client := &http.Client{}
		req, err := http.NewRequest(method, slackWebhook, strings.NewReader(body))
		if err != nil {
			log.Println("Http Error: ", err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			log.Println("Http:", err)
		}

		fmt.Println(resp)

	}

	return nil
}
