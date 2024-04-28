locals {
  account_id        = data.aws_caller_identity.current.account_id
  aws_region        = data.aws_region.current.name
  vpc_id            = var.vpc_id
  first_subnet_id   = var.first_subnet_id
  second_subnet_id  = var.second_subnet_id
  security_group_id = var.security_group_id
}
