# Terraform Plan Validator
Validates Terraform Plans using TFSEC and OPA

# Commands
go run main.go check "delete-rg-test.json" "azure"

# Docker
``` 
docker build . -t terraform-plan-validator
docker tag terraform-plan-validator bradmccoydev/terraform-plan-validator:latest
docker push bradmccoydev/terraform-plan-validator:latest
docker pull bradmccoydev/terraform-plan-validator:latest

docker run -p 80:80 bradmccoydev/terraform-plan-validator:latest check "delete-rg-test.json" "azure"
```

variables:
    OPA_GCP_POLICY: opa-gcp-policy.rego
    OPA_AZURE_POLICY: opa-azure-policy.rego
    OPA_AWS_POLICY: opa-aws-policy.rego
    OPA_REGO_QUERY: data.terraform.analysis.authz

chmod +x ./main
./main check "delete-rg-test.json" "azure"
