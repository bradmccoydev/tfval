# Terraform Plan Validator
Validates Terraform Plans using TFSEC and OPA

# Commands
go run main.go check "delete-rg-test.json" "azure"

# Docker
``` 
docker build . -t terrform-plan-validator
docker tag terrform-plan-validator bradmccoydev/terrform-plan-validator:latest
docker push bradmccoydev/terrform-plan-validator:latest
docker pull bradmccoydev/terrform-plan-validator:latest
```
