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

resource "aws_iam_policy" "landing_point_policy" {
  name        = "landing_point_policy"
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
          "arn:aws:s3:::landing_point"
        ]
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "attach_landing_point_policy_to_worker" {
  role       = aws_iam_role.worker_role.name
  policy_arn = aws_iam_policy.landing_point_policy.arn
}
