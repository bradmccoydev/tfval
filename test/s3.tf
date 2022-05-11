resource "aws_s3_bucket" "b" {
  bucket = "terraform-plan-validator"

  tags = {
    Name        = "terraform-plan-validator"
    Environment = "Dev"
  }
}

# terraform init && terraform plan -out test.tfplan