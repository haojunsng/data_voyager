resource "aws_iam_policy" "gomu_landing_bucket_policy" {
  name        = "gomu_landing_bucket_policy"
  path        = "/"
  description = "Allow "

  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "BucketPolicy",
        "Effect" : "Allow",
        "Action" : [
          "s3:PutObject",
          "s3:GetObject",
          "s3:ListBucket",
          "s3:DeleteObject"
        ],
        "Resource" : [
          "arn:aws:s3:::*/*",
          "arn:aws:s3:::gomu_landing_bucket"
        ]
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "attach_gomu_landing_bucket_policy_to_worker" {
  role       = aws_iam_role.ecs_task_role.name
  policy_arn = aws_iam_policy.gomu_landing_bucket_policy.arn
}

# To attach to Airflow Role
resource "aws_iam_policy" "airflow_policy" {
  name        = "gomu_landing_bucket_policy"
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
      }
    ]
  })
}
