package terraform.analysis

import input as tfplan

########################
# Parameters for Policy
########################
critical_resource_score = { "delete" : 99, "create" : 0, "modify" : 99 }
important_resource_score = { "delete" : 99, "create" : 0, "modify" : 1 }
lowrisk_resource_score = { "delete" : 1, "create" : 0, "modify" : 0 }

# acceptable score for automated authorization
max_acceptable_score = 2

# weights assigned for each operation on each resource-type
weights = {
    # CRITICAL RESOURCES - Any changes require human cross check
    "azurerm_mssql_server": critical_resource_score,
    "azurerm_mssql_database": critical_resource_score,
    "azurerm_sql_failover_group": critical_resource_score,
    "azurerm_servicebus_namespace": critical_resource_score,
    "azurerm_servicebus_queue": critical_resource_score,
    "azurerm_servicebus_topic": critical_resource_score,
    "azurerm_servicebus_subscription": critical_resource_score,
    "azurerm_virtual_network": critical_resource_score,
    "azurerm_subnet": critical_resource_score,
    "azurerm_key_vault_secret": critical_resource_score,

    # IMPORTANT RESOURCES - Minor modifications acceptable. Deletion prohibited.
    "azurerm_app_service_plan": important_resource_score,
    "azurerm_container_registry": important_resource_score,
    "azurerm_data_factory": important_resource_score,
    "azurerm_databricks_workspace": important_resource_score,
    "azurerm_key_vault": important_resource_score,
    "azurerm_notification_hub_namespace": important_resource_score,
    "azurerm_resource_group": important_resource_score,
    "azurerm_storage_account": important_resource_score,
    "azurerm_nat_gateway": important_resource_score,
    "azurerm_kubernetes_cluster": important_resource_score,
    "azurerm_kubernetes_cluster_node_pool": important_resource_score,

    # LOW RISK - Generally allowed.
    "azurerm_application_insights": lowrisk_resource_score,
    "azurerm_app_service": lowrisk_resource_score,
    "azurerm_app_service_slot": lowrisk_resource_score,
    "azurerm_function_app": lowrisk_resource_score,
    "azurerm_function_app_slot": lowrisk_resource_score,
    "azurerm_key_vault_access_policy": lowrisk_resource_score,
    "azurerm_log_analytics_workspace": lowrisk_resource_score,
    "azurerm_storage_container": lowrisk_resource_score,
}

# Consider exactly these resource types in calculations
resource_types = {
    # CRITICAL RESOURCES
    "azurerm_mssql_server",
    "azurerm_mssql_database",
    "azurerm_sql_failover_group",
    "azurerm_servicebus_namespace",
    "azurerm_servicebus_queue",
    "azurerm_servicebus_topic",
    "azurerm_servicebus_subscription",
    "azurerm_virtual_network",
    "azurerm_subnet",
    "azurerm_key_vault_secret",

    # IMPORTANT RESOURCES
    "azurerm_app_service_plan",
    "azurerm_container_registry",
    "azurerm_data_factory",
    "azurerm_databricks_workspace",
    "azurerm_key_vault",
    "azurerm_notification_hub_namespace",
    "azurerm_resource_group",
    "azurerm_storage_account",
    "azurerm_nat_gateway",
    "azurerm_kubernetes_cluster",
    "azurerm_kubernetes_cluster_node_pool",

    # LOW RISK RESOURCES
    "azurerm_application_insights",
    "azurerm_app_service",
    "azurerm_app_service_slot",
    "azurerm_function_app",
    "azurerm_function_app_slot",
    "azurerm_key_vault_access_policy",
    "azurerm_log_analytics_workspace",
    "azurerm_storage_container",
}

#########
# Policy
#########

# Authorization holds if score for the plan is acceptable and no changes are made to IAM
default authz = false
authz {
    score < max_acceptable_score
    not touches_iam
}

# Compute the score for a Terraform plan as the weighted sum of deletions, creations, modifications
score = s {
    all := [ x |
            some resource_type
            crud := weights[resource_type];
            del := crud["delete"] * num_deletes[resource_type];
            new := crud["create"] * num_creates[resource_type];
            mod := crud["modify"] * num_modifies[resource_type];
            x := del + new + mod
    ]
    s := sum(all)
}

# Whether there is any change to IAM
touches_iam {
    all := resources["aws_iam"]
    count(all) > 0
}

####################
# Terraform Library
####################

# list of all resources of a given type
resources[resource_type] = all {
    some resource_type
    resource_types[resource_type]
    all := [name |
        name:= tfplan.resource_changes[_]
        name.type == resource_type
    ]
}

# number of creations of resources of a given type
num_creates[resource_type] = num {
    some resource_type
    resource_types[resource_type]
    all := resources[resource_type]
    creates := [res |  res:= all[_]; res.change.actions[_] == "create"]
    num := count(creates)
}


# number of deletions of resources of a given type
num_deletes[resource_type] = num {
    some resource_type
    resource_types[resource_type]
    all := resources[resource_type]
    deletions := [res |  res:= all[_]; res.change.actions[_] == "delete"]
    num := count(deletions)
}

# number of modifications to resources of a given type
num_modifies[resource_type] = num {
    some resource_type
    resource_types[resource_type]
    all := resources[resource_type]
    modifies := [res |  res:= all[_]; res.change.actions[_] == "update"]
    num := count(modifies)
}

deny[msg] {
    msg := sprintf("{\"validation_passed\":%s,\"score\":%d,\"max_acceptable_score\":%d,\"weights\":[%s], \"data\": %s },", [
        score < max_acceptable_score,
        score,
        max_acceptable_score,
        weights,
        input.resource_changes[_]
    ])
}
