# Create CD Github Action service user
resource "aws_iam_user" "gha_service_user" {
  name = "gha_service_user"
}

resource "aws_iam_user_policy_attachment" "gha_service_user_ecr_access" {
  user       = aws_iam_user.gha_service_user.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryFullAccess"
}
