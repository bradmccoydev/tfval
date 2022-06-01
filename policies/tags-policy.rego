package terraform.common

deny[msg] {
    changeset := input.resource_changes[_]
    
    required_tags := {"Environment", "Owner","x"}
    provided_tags := {tag | changeset.change.after.tags_all[tag]}
    missing_tags := required_tags - provided_tags
    
    count(missing_tags) > 0
    
    msg := sprintf("Resource %v missing tags: %v", [
        changeset.address,
        concat(", ", missing_tags)
    ])
}