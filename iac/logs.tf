resource "aws_cloudwatch_log_group" "gomu_log_group" {
  name = "gomu-log-group"
}

resource "aws_cloudwatch_log_stream" "gomu_log_stream" {
  name           = "gomu-log-stream"
  log_group_name = aws_cloudwatch_log_group.gomu_log_group.name
}

resource "aws_cloudwatch_log_group" "kafka_log_group" {
  name = "kafka-broker-logs"
}
