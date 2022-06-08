package tfsec

import (
	"fmt"
	"testing"
)

func TestGetTfScore(t *testing.T) {
	testCases := []struct {
		name     string
		tfReport string
		err      error
	}{
		{"critical-severity", "{\"results\": [{\"rule_id\": \"azure-keyvault-specify-network-acl\",\"legacy_rule_id\": \"AZU020\",\"rule_description\": \"Key vault should have the network acl block specified\",\"rule_provider\": \"azure\",\"impact\": \"Without a network ACL the key vault is freely accessible\",\"resolution\": \"Set a network ACL for the key vault\",\"links\": [\"https://tfsec.dev/docs/azure/keyvault/specify-network-acl#azure/keyvault\",\"https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/key_vault#network_acls\",\"https://docs.microsoft.com/en-us/azure/key-vault/general/network-security\"],\"description\": \"Resource 'azurerm_key_vault.default' specifies does not specify a network acl block with default action.\",\"severity\": \"CRITICAL\",\"status\": \"failed\",\"location\": {\"filename\": \"/opt/atlassian/pipelines/agent/build/azurerm/azurerm_key_vault/main.tf\",\"start_line\": 1,\"end_line\": 13}}]}", nil},
		{"medium-severity", "{\"results\": [{\"rule_id\": \"azure-database-enable-audit\",\"legacy_rule_id\": \"AZU018\",\"rule_description\": \"Auditing should be enabled on Azure SQL Databases\",\"rule_provider\": \"azure\",\"impact\": \"Auditing provides valuable information about access and usage\",\"resolution\": \"Enable auditing on Azure SQL databases\",\"links\": [\"https://tfsec.dev/docs/azure/database/enable-audit#azure/database\",\"https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/sql_server#extended_auditing_policy\",\"https://docs.microsoft.com/en-us/azure/azure-sql/database/auditing-overview\"],\"description\": \"Resource 'azurerm_mssql_server.default' does not have an extended audit policy configured.\",\"severity\": \"MEDIUM\",\"status\": \"failed\",\"location\": {\"filename\": \"/opt/atlassian/pipelines/agent/build/azurerm/azurerm_mssql_server/main.tf\",\"start_line\": 1,\"end_line\": 11}}]}", nil},
	}
	for _, tc := range testCases {
		tfSecReport := CheckIfPlanPassesTfPolicy([]byte(tc.tfReport), "CRITICAL")
		fmt.Println(tc.name, tfSecReport)
	}
}
