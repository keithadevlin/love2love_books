provider "aws" {
  version = "1.60"
  region  = "eu-west-1"
}

terraform {
  backend "s3" {
    bucket = "love2love-tf-state"
    key    = "mws_sales_report/terraform.tfstate"
    region = "eu-west-1"
  }
}

module "mws_sales_report" {
  source      = "../../../modules/mws_sales_report"
  environment = "prod"
}
