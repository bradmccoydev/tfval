# Terraform Plan Validator
Validates Terraform Plans using TFSEC and OPA

# Docker
``` 
docker build . -t terrform-plan-validator
docker tag terrform-plan-validator bradmccoydev/terrform-plan-validator:latest
docker push terrform-plan-validator/terrform-plan-validator:latest
docker pull bradmccoydev/terrform-plan-validator:latest
```