![passing](https://github.com/bradmccoydev/terraform-plan-validator/actions/workflows/ci.yml/badge.svg)
# Terraform Plan Validator
Validates Terraform Plans using TFSEC and OPA

# Commands
go run main.go check "delete-rg-test.json" "azure"
env GOOS=linux GOARCH=amd64 go build -o terraform-plan-validator-amd64

# Docker
``` 
docker build . -t terraform-plan-validator
docker tag terraform-plan-validator bradmccoydev/terraform-plan-validator:latest
docker push bradmccoydev/terraform-plan-validator:latest
docker pull bradmccoydev/terraform-plan-validator:latest

docker run -p 80:80 bradmccoydev/terraform-plan-validator:latest check "delete-rg-test.json" "azure"
```

### Variables

| Variable | Value |
| --- | --- |
| OPA_GCP_POLICY | opa-gcp-policy.rego |
| OPA_AZURE_POLICY | opa-azure-policy.rego |
| OPA_AWS_POLICY | opa-aws-policy.rego |
| OPA_REGO_QUERY | data.terraform.analysis.authz |
| TFSEC_MAX_SEVERITY | ["LOW", "MEDIUM", "CRITICAL"] |