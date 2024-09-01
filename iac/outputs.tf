output "msk_bootstrap_servers_tls" {
  value = aws_msk_cluster.kafka.bootstrap_brokers_tls
}
