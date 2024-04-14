# Airflow MWAA role
resource "aws_iam_role" "airflow_role" {
  name               = "gomu_airflow_role"
  assume_role_policy = data.aws_iam_policy_document.mwaa_assume.json
}

# Airflow MWAA assume role policy
data "aws_iam_policy_document" "mwaa_assume" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["airflow.amazonaws.com"]
    }

    principals {
      type        = "Service"
      identifiers = ["airflow-env.amazonaws.com"]
    }

    principals {
      type        = "Service"
      identifiers = ["batch.amazonaws.com"]
    }

    principals {
      type        = "Service"
      identifiers = ["ssm.amazonaws.com"]
    }
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
    principals {
      type        = "Service"
      identifiers = ["s3.amazonaws.com"]
    }
  }
}

# Airflow MWAA role policy
resource "aws_iam_policy" "airflow_policy" {
  name        = "airflow_policy"
  path        = "/"
  description = "Allow "

  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "ECSPolicy",
        "Effect" : "Allow",
        "Action" : [
          "ecs:RunTask",
          "ecs:DescribeTasks"
        ],
        "Resource" : [aws_ecs_task_definition.definition.arn]
      },
      {
        "Sid" : "PassRole",
        "Effect" : "Allow",
        "Action" : [
          "iam:PassRole"
        ],
        "Resource" : [aws_iam_role.ecs_task_execution_role.arn]
      },
      {
        "Sid" : "ECSTaskDescription",
        "Effect" : "Allow",
        "Action" : [
          "ecs:DescribeTasks"
        ],
        "Resource" : ["arn:aws:ecs:${local.aws_region}:${local.account_id}:task/${aws_ecs_cluster.ecs_cluster.name}/*"]
      },
      {
        "Sid" : "LogsFlow",
        "Effect" : "Allow",
        "Action" : [
          "logs:CreateLogStream",
          "logs:CreateLogGroup",
          "logs:PutLogEvents",
          "logs:GetLogEvents",
          "logs:GetLogRecord",
          "logs:GetLogGroupFields",
          "logs:GetQueryResults"
        ],
        "Resource" : ["arn:aws:logs:${local.aws_region}:${local.account_id}:log-group:${aws_cloudwatch_log_group.gomu_log_group.name}:log-stream:${aws_cloudwatch_log_stream.gomu_log_stream.name}/*"]
      },
      {
        "Sid" : "S3",
        "Effect" : "Allow",
        "Action" : [
          "s3:*"
        ],
        "Resource" : [
          aws_s3_bucket.gomu_airflow_s3.arn,
          "${aws_s3_bucket.gomu_airflow_s3.arn}/*"
        ]
      }
    ]
  })
}

resource "aws_mwaa_environment" "gomu_airflow" {
  count              = var.deploy_mwaa ? 1 : 0
  dag_s3_path        = "dags/"
  execution_role_arn = aws_iam_role.airflow_role.arn
  name               = "gomu_airflow"

  network_configuration {
    security_group_ids = [var.security_group_id]
    subnet_ids         = [var.first_subnet_id, var.second_subnet_id]
  }

  source_bucket_arn = aws_s3_bucket.gomu_airflow_s3.arn
}

resource "aws_s3_bucket" "gomu_airflow_s3" {
  bucket = "gomu-airflow-s3"
}

resource "aws_s3_bucket_versioning" "this" {
  bucket = aws_s3_bucket.gomu_airflow_s3.id
  versioning_configuration {
    status = "Enabled"
  }
}
resource "aws_s3_bucket_server_side_encryption_configuration" "this" {
  bucket = aws_s3_bucket.gomu_airflow_s3.id

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

resource "aws_s3_bucket_public_access_block" "this" {
  bucket                  = aws_s3_bucket.gomu_airflow_s3.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}
