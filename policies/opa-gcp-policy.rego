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
    "google_container_cluster": critical_resource_score,

    # IMPORTANT RESOURCES - Minor modifications acceptable. Deletion prohibited.
    "google_container_node_pool": important_resource_score,

    # LOW RISK - Generally allowed.
    "google_pubsub_subscription": lowrisk_resource_score,
}


# Consider exactly these resource types in calculations
resource_types = {
    # CRITICAL RESOURCES
    "google_container_cluster",

    # IMPORTANT RESOURCES
    "google_container_node_pool",

    # LOW RISK RESOURCES
    "google_pubsub_subscription",
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
