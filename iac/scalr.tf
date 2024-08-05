resource "aws_iam_openid_connect_provider" "scalr" {
  url             = "https://scalr.io"
  client_id_list  = ["roronoa"]
  thumbprint_list = ["08745487e891c19e3078c1f2a07e452950ef36f6"]
}

data "aws_iam_policy_document" "scalr_assume_role" {
  statement {
    actions = ["sts:AssumeRoleWithWebIdentity"]
    principals {
      type        = "Federated"
      identifiers = [aws_iam_openid_connect_provider.scalr.arn]
    }
    condition {
      test     = "StringLike"
      variable = "scalr.io:sub"
      values   = ["account:roronoa:environment:Environment-A:workspace:Roronoa"]
    }
  }
}

resource "aws_iam_role" "scalr" {
  name               = "scalr-roronoa-zoro"
  assume_role_policy = data.aws_iam_policy_document.scalr_assume_role.json
}

resource "aws_iam_role_policy_attachment" "scalr" {
  role       = aws_iam_role.scalr.name
  policy_arn = "arn:aws:iam::aws:policy/AdministratorAccess"
}
