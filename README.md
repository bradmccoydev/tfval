![passing](https://github.com/bradmccoydev/terraform-plan-validator/actions/workflows/ci.yml/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/bradmccoydev/terraform-plan-validator)](https://goreportcard.com/report/github.com/bradmccoydev/terraform-plan-validator) ![GitHub](https://img.shields.io/github/license/bradmccoydev/terraform-plan-validator) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/bradmccoydev/terraform-plan-validator)
# Terraform Plan Validator
Validates Terraform Plans using TFSEC and OPA

# Commands
go run main.go check --planFileName "delete-rg-test.json" --CloudProvider "azure"

go run main.go opascore --planFileName "delete-rg-test.json" --cloudProvider "azure"


env GOOS=linux GOARCH=amd64 go build -o terraform-plan-validator-amd64

# Docker
``` 
docker pull bradmccoydev/terraform-plan-validator:latest
```

```
docker run -p 80:80 bradmccoydev/terraform-plan-validator:latest check --PlanFileName "delete-rg-test.json" --CloudProvider "azure"
```

```
docker run -p 80:80 bradmccoydev/terraform-plan-validator:latest sendreport --prNumber "$test" --repoFullUrl "http://www.google.com" --fileName "./mock.json" --slackWebhook "webhook"
```

### Variables

| Variable | Value |
| --- | --- |
| OPA_GCP_POLICY | opa-gcp-policy.rego |
| OPA_AZURE_POLICY | opa-azure-policy.rego |
| OPA_AWS_POLICY | opa-aws-policy.rego |
| OPA_REGO_QUERY | data.terraform.analysis.authz |
| TFSEC_MAX_SEVERITY | ["LOW", "MEDIUM", "CRITICAL"] |

Maintainers:
* Brad McCoy ([@bradmccoydev](https://github.com/bradmccoydev)), Moula
* Ben Poh ([@benhpoh](https://github.com/benhpoh)), Moula

# License

Terraform Plan Validator is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/bradmccoydev/terraform-plan-validator/blob/main/LICENSE)
