package tfsec

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/bradmccoydev/tfval/model"
)

var (
	hasItems bool
	body     string
)

func TestGetTfSecSlack(t *testing.T) {
	testCases := []struct {
		slackWebhook     string
		prNumber         string
		repoFullUrl      string
		fileName         string
		tfsecMaxSeverity string
		err              error
	}{
		{"addme", "1", "https://github.com/bradmccoydev/tfval", "mock.json", "{\"results\": [{\"rule_id\": \"azure-keyvault-specify-network-acl\",\"legacy_rule_id\": \"AZU020\",\"rule_description\": \"Key vault should have the network acl block specified\",\"rule_provider\": \"azure\",\"impact\": \"Without a network ACL the key vault is freely accessible\",\"resolution\": \"Set a network ACL for the key vault\",\"links\": [\"https://tfsec.dev/docs/azure/keyvault/specify-network-acl#azure/keyvault\",\"https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/key_vault#network_acls\",\"https://docs.microsoft.com/en-us/azure/key-vault/general/network-security\"],\"description\": \"Resource 'azurerm_key_vault.default' specifies does not specify a network acl block with default action.\",\"severity\": \"CRITICAL\",\"status\": \"failed\",\"location\": {\"filename\": \"/opt/atlassian/pipelines/agent/build/azurerm/azurerm_key_vault/main.tf\",\"start_line\": 1,\"end_line\": 13}}]}", nil},
		{"addme", "1", "https://github.com/bradmccoydev/tfval", "mock.json", "{\"results\": [{\"rule_id\": \"azure-database-enable-audit\",\"legacy_rule_id\": \"AZU018\",\"rule_description\": \"Auditing should be enabled on Azure SQL Databases\",\"rule_provider\": \"azure\",\"impact\": \"Auditing provides valuable information about access and usage\",\"resolution\": \"Enable auditing on Azure SQL databases\",\"links\": [\"https://tfsec.dev/docs/azure/database/enable-audit#azure/database\",\"https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/sql_server#extended_auditing_policy\",\"https://docs.microsoft.com/en-us/azure/azure-sql/database/auditing-overview\"],\"description\": \"Resource 'azurerm_mssql_server.default' does not have an extended audit policy configured.\",\"severity\": \"MEDIUM\",\"status\": \"failed\",\"location\": {\"filename\": \"/opt/atlassian/pipelines/agent/build/azurerm/azurerm_mssql_server/main.tf\",\"start_line\": 1,\"end_line\": 11}}]}", nil},
	}

	for _, tc := range testCases {
		report, err := ioutil.ReadFile(tc.fileName)
		if err != nil {
			fmt.Println(err)
		}

		tc.tfsecMaxSeverity = "CRITICAL"

		var vulnerabilities model.Vulnerabilities
		json.Unmarshal([]byte(report), &vulnerabilities)

		if len(vulnerabilities.Results) > 0 {
			header := fmt.Sprintf(`{"blocks": [{"type": "header","text": {"type": "plain_text","text": ":cop: Pull Request %v Static Code Analysis Failed :cop:","emoji": true}}`, tc.prNumber)
			footer := fmt.Sprintf(`,{"type": "divider"},{"type": "section","text": {"type": "mrkdwn","text": "View further details in the pull request:"},"accessory": {"type": "button","text": {"type": "plain_text","text": "View Pull Request","emoji": true},"value": "click_me_123","url": "%v","action_id": "button-action"}}`, tc.repoFullUrl)
			body = fmt.Sprintf("%v%v", header, footer)

			for _, element := range vulnerabilities.Results {
				if element.Severity == "LOW" && (tc.tfsecMaxSeverity == "LOW" || tc.tfsecMaxSeverity == "MEDIUM") {
					produceSlackBlockLineItem(element.Impact, element.Resolution)
				} else if element.Severity == "MEDIUM" && tc.tfsecMaxSeverity == "MEDIUM" {
					produceSlackBlockLineItem(element.Impact, element.Resolution)
				} else if element.Severity == "CRITICAL" && tc.tfsecMaxSeverity == "CRITICAL" {
					produceSlackBlockLineItem(element.Impact, element.Resolution)
				}
			}

			if hasItems {
				client := &http.Client{}
				body = fmt.Sprintf(`%v%v`, body, "]}")
				req, err := http.NewRequest("POST", tc.slackWebhook, strings.NewReader(body))
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
	}
}

func produceSlackBlockLineItem(impact string, resolution string) {
	hasItems = true
	impactLine := fmt.Sprintf(`,{"type": "section","text": {"type": "mrkdwn","text": "%v *%v*"}}`, ":orange_book:", impact)
	resolutionLine := fmt.Sprintf(`,{"type": "section","text": {"type": "mrkdwn","text": ":arrow_right: %v"}}`, resolution)
	body = fmt.Sprintf(`%v%v%v`, body, impactLine, resolutionLine)
}
