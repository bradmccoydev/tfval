resource "aws_s3_bucket" "a" {
  bucket = "tfval-brad"

  tags = {
    Owner       = "bradmccoy"
    SourcePath  = "https://github.com/bradmccoydev/tfval"
    Environment = "Dev"
    Provisioner = "Terraform"
  }
}

resource "aws_s3_bucket" "b" {
  bucket = "tfval"

  tags = {
    Name        = "tfval"
    Environment = "Dev"
  }
}

# terraform init && terraform plan -out test.tfplan
