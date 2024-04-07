# `IaC/`

## Description

### Resources provisioned/referenced through Terraform include:
---
- Networking
    - VPC
    - Subnets
    - Security Group
- ECS Task Definition
- ECR
- Identity Access Management
    - Service User for GHA
    - ECS Task Execution Role
    - ECS Task Role
- Cloudwatch Logs
- S3 Buckets

### Resources NOT provisioned/referenced through Terraform
---
- SSM Parameters

Only supports local `terraform init/plan/apply`, `.tfstate` files are maintained locally -- does not support collaboration (Scalr is not free).
