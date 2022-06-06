package model

type tfValidation []struct {
	ValidationPassed bool `json:"validation_passed"`
	Weights          []struct {
		AzurermAppService struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_app_service"`
		AzurermAppServicePlan struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_app_service_plan"`
		AzurermAppServiceSlot struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_app_service_slot"`
		AzurermApplicationInsights struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_application_insights"`
		AzurermContainerRegistry struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_container_registry"`
		AzurermDataFactory struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_data_factory"`
		AzurermDatabricksWorkspace struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_databricks_workspace"`
		AzurermFunctionApp struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_function_app"`
		AzurermFunctionAppSlot struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_function_app_slot"`
		AzurermKeyVault struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_key_vault"`
		AzurermKeyVaultAccessPolicy struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_key_vault_access_policy"`
		AzurermKeyVaultSecret struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_key_vault_secret"`
		AzurermKubernetesCluster struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_kubernetes_cluster"`
		AzurermKubernetesClusterNodePool struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_kubernetes_cluster_node_pool"`
		AzurermLogAnalyticsWorkspace struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_log_analytics_workspace"`
		AzurermMssqlDatabase struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_mssql_database"`
		AzurermMssqlServer struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_mssql_server"`
		AzurermNatGateway struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_nat_gateway"`
		AzurermNotificationHubNamespace struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_notification_hub_namespace"`
		AzurermResourceGroup struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_resource_group"`
		AzurermServicebusNamespace struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_servicebus_namespace"`
		AzurermServicebusQueue struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_servicebus_queue"`
		AzurermServicebusSubscription struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_servicebus_subscription"`
		AzurermServicebusTopic struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_servicebus_topic"`
		AzurermSQLFailoverGroup struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_sql_failover_group"`
		AzurermStorageAccount struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_storage_account"`
		AzurermStorageContainer struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_storage_container"`
		AzurermSubnet struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_subnet"`
		AzurermVirtualNetwork struct {
			Create int `json:"create"`
			Delete int `json:"delete"`
			Modify int `json:"modify"`
		} `json:"azurerm_virtual_network"`
	} `json:"weights"`
	Data struct {
		Address string `json:"address"`
		Change  struct {
			Actions        []string    `json:"actions"`
			After          interface{} `json:"after"`
			AfterSensitive bool        `json:"after_sensitive"`
			AfterUnknown   struct {
			} `json:"after_unknown"`
			Before struct {
				ID       string `json:"id"`
				Location string `json:"location"`
				Name     string `json:"name"`
				Tags     struct {
				} `json:"tags"`
				Timeouts interface{} `json:"timeouts"`
			} `json:"before"`
			BeforeSensitive struct {
				Tags struct {
				} `json:"tags"`
			} `json:"before_sensitive"`
		} `json:"change"`
		Mode         string `json:"mode"`
		Name         string `json:"name"`
		ProviderName string `json:"provider_name"`
		Type         string `json:"type"`
	} `json:"data"`
}
