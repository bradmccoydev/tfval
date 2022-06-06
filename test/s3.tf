resource "aws_s3_bucket" "b" {
  bucket = "tfval"

  tags = {
    Name        = "tfval"
    Environment = "Dev"
  }
}

# terraform init && terraform plan -out test.tfplan