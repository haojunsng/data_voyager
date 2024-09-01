resource "aws_msk_configuration" "kafka_config" {
  kafka_versions    = [local.kafka_versions]
  name              = "${local.kafka_name}-config"
  server_properties = <<EOF
auto.create.topics.enable = true
delete.topic.enable = true
EOF
}

resource "aws_msk_cluster" "kafka" {
  cluster_name           = local.kafka_name
  kafka_version          = local.kafka_versions
  number_of_broker_nodes = local.kafka_number_of_broker_nodes
  broker_node_group_info {
    instance_type = local.kafka_instance_size
    storage_info {
      ebs_storage_info {
        volume_size = local.kafka_ebs_storage_volume_size
      }
    }
    client_subnets  = [var.first_subnet_id, var.second_subnet_id]
    security_groups = [var.security_group_id]
  }
  encryption_info {
    encryption_in_transit {
      client_broker = local.kafka_in_transit_encryption
    }
    encryption_at_rest_kms_key_arn = aws_kms_key.kafka_kms_key.arn
  }
  configuration_info {
    arn      = aws_msk_configuration.kafka_config.arn
    revision = aws_msk_configuration.kafka_config.latest_revision
  }
  logging_info {
    broker_logs {
      cloudwatch_logs {
        enabled   = true
        log_group = aws_cloudwatch_log_group.kafka_log_group.name
      }
    }
  }
}
