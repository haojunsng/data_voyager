terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "> 3.27"
    }
  }
  backend "remote" {
    hostname     = "roronoa.scalr.io"
    organization = "env-v0od6uicoc86gsnsf"

    workspaces {
      name = "Roronoa"
    }
  }
}

provider "aws" {
  region = "ap-southeast-1"
}

provider "kubernetes" {
  config_context = module.eks.cluster_arn
}
