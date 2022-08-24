package terraform.common

resources_types = [
	"aws_neptune_cluster",
]

max_acceptable_score = 0

deny[msg] {
	changeset := input.resource_changes[_]

	array_contains(resources_types, changeset.type)

	msg := sprintf("{\"validation_passed\":%s,\"score\":%d,\"max_acceptable_score\":%d,\"data\":{\"address\": \"%s\"}},", [
		"false",
		count(resources_types),
		max_acceptable_score,
		changeset.address,
	])
}

array_contains(arr, elem) {
	arr[_] = elem
}
