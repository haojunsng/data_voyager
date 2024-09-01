locals {
  account_id        = data.aws_caller_identity.current.account_id
  aws_region        = data.aws_region.current.name
  vpc_id            = var.vpc_id
  first_subnet_id   = var.first_subnet_id
  second_subnet_id  = var.second_subnet_id
  security_group_id = var.security_group_id

  # kafka
  kafka_name                    = "tenryuubito"
  kafka_versions                = "3.5.1"
  kafka_instance_size           = "kafka.t3.small"
  kafka_number_of_broker_nodes  = 2
  kafka_ebs_storage_volume_size = 100
  kafka_in_transit_encryption   = "PLAINTEXT"
}
