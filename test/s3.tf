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
    Owner       = "bradmccoy"
    SourcePath  = "https://github.com/bradmccoydev/tfval"
    Environment = "Dev"
    Provisioner = "Terraform"
  }
}

resource "aws_neptune_cluster" "default" {
  cluster_identifier                  = "neptune-cluster-demo"
  engine                              = "neptune"
  backup_retention_period             = 5
  preferred_backup_window             = "07:00-09:00"
  skip_final_snapshot                 = true
  iam_database_authentication_enabled = true
  apply_immediately                   = true
}

# terraform init && terraform plan -out test.tfplan
