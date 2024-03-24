resource "aws_iam_role" "worker_role" {
  name = "worker_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      },
    ]
  })
}

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
  role       = aws_iam_role.worker_role.name
  policy_arn = aws_iam_policy.gomu_landing_bucket_policy.arn
}
