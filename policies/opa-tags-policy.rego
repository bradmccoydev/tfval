package terraform.common

resources_types = [
	"aws_s3_bucket",
	"aws_dynamodb_table",
]

# acceptable score for automated authorization
max_acceptable_score = 0

deny[msg] {
	changeset := input.resource_changes[_]

	required_tags := {
		"Owner",
		"SourcePath",
		"Environment",
		"Provisioner",
	}

	provided_tags := {tag | changeset.change.after.tags_all[tag]}
	missing_tags := required_tags - provided_tags

	#array_contains(resources_types, changeset.type)

	count(missing_tags) > 0

	msg := sprintf("{\"validation_passed\":%s,\"score\":%d,\"max_acceptable_score\":%d,\"data\":{\"address\": \"%s\"}},", [
		count(missing_tags) <= max_acceptable_score,
		count(missing_tags),
		max_acceptable_score,
		changeset.address,
	])
}

array_contains(arr, elem) {
	arr[_] = elem
}
