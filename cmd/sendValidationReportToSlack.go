package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/bradmccoydev/tfval/model"
	"github.com/spf13/cobra"
)

var (
	hasItems     bool
	body         string
	prNumber     string
	repoFullUrl  string
	fileName     string
	slackWebhook string

	sendValidationReportToSlackCmd = &cobra.Command{
		Use:   "sendreport",
		Short: "Send Validation Report",
		Long:  `Send Terraform validation Report to slack`,
		Run: func(cmd *cobra.Command, args []string) {
			err := sendValidationReportToSlack(args)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(sendValidationReportToSlackCmd)
	sendValidationReportToSlackCmd.PersistentFlags().StringVarP(&prNumber, "prNumber", "p", prNumber, "Pull request number")
	sendValidationReportToSlackCmd.PersistentFlags().StringVarP(&repoFullUrl, "repoFullUrl", "r", repoFullUrl, "Full repo URL")
	sendValidationReportToSlackCmd.PersistentFlags().StringVarP(&fileName, "fileName", "f", fileName, "Filename of the tfsec report")
	sendValidationReportToSlackCmd.PersistentFlags().StringVarP(&slackWebhook, "slackWebhook", "s", slackWebhook, "The Slack Webhook")
	sendValidationReportToSlackCmd.PersistentFlags().StringVarP(&tfsecMaxSeverity, "tfsecMaxSeverity", "t", tfsecMaxSeverity, "The TF Sec Max Severity")
}

func sendValidationReportToSlack(args []string) error {
	report, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	var vulnerabilities model.Vulnerabilities
	json.Unmarshal([]byte(report), &vulnerabilities)

	if len(vulnerabilities.Results) > 0 {
		header := fmt.Sprintf(`{"blocks": [{"type": "header","text": {"type": "plain_text","text": ":cop: Pull Request %v Static Code Analysis Failed :cop:","emoji": true}}`, prNumber)
		footer := fmt.Sprintf(`,{"type": "divider"},{"type": "section","text": {"type": "mrkdwn","text": "View further details in the pull request:"},"accessory": {"type": "button","text": {"type": "plain_text","text": "View Pull Request","emoji": true},"value": "click_me_123","url": "%v","action_id": "button-action"}}`, repoFullUrl)
		body = fmt.Sprintf("%v%v", header, footer)

		for _, element := range vulnerabilities.Results {
			if element.Severity == "LOW" && (tfsecMaxSeverity == "LOW" || tfsecMaxSeverity == "MEDIUM") {
				produceSlackBlockLineItem(element.Impact, element.Resolution)
			} else if element.Severity == "MEDIUM" && tfsecMaxSeverity == "MEDIUM" {
				produceSlackBlockLineItem(element.Impact, element.Resolution)
			} else if element.Severity == "CRITICAL" && tfsecMaxSeverity == "CRITICAL" {
				produceSlackBlockLineItem(element.Impact, element.Resolution)
			}
		}

		if hasItems {
			client := &http.Client{}
			body = fmt.Sprintf(`%v%v`, body, "]}")
			req, err := http.NewRequest("POST", slackWebhook, strings.NewReader(body))
			if err != nil {
				log.Println("Http Error: ", err)
			}

			req.Header.Set("Content-Type", "application/json")
			resp, err := client.Do(req)
			if err != nil {
				log.Println("Http Error: ", err)
			}

			fmt.Println(resp)
		}
	}

	return nil
}

func produceSlackBlockLineItem(impact string, resolution string) {
	hasItems = true
	impactLine := fmt.Sprintf(`,{"type": "section","text": {"type": "mrkdwn","text": "%v *%v*"}}`, ":orange_book:", impact)
	resolutionLine := fmt.Sprintf(`,{"type": "section","text": {"type": "mrkdwn","text": ":arrow_right: %v"}}`, resolution)
	body = fmt.Sprintf(`%v%v%v`, body, impactLine, resolutionLine)
}
