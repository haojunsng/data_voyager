module "eks" {
  source          = "terraform-aws-modules/eks/aws"
  cluster_name    = "jinbei"
  cluster_version = "1.29"
  subnet_ids      = [var.first_subnet_id, var.second_subnet_id]

  vpc_id = var.vpc_id

  eks_managed_node_group_defaults = {
    ami_type = "AL2_x86_64"

  }

  eks_managed_node_groups = {
    one = {
      name = "node-group-1"

      instance_types = ["t3.small"]

      min_size     = 1
      max_size     = 3
      desired_size = 2
    }

    two = {
      name = "node-group-2"

      instance_types = ["t3.small"]

      min_size     = 1
      max_size     = 2
      desired_size = 1
    }
  }
}

data "aws_ssm_parameter" "kafka-broker" {
  name = "kafka-broker"
}

resource "kubernetes_secret" "kafka-broker" {
  metadata {
    name      = "kafka-broker"
    namespace = "default"
  }
  data = {
    password = data.aws_ssm_parameter.kafka-broker.value
  }
}

resource "aws_iam_role" "irsa_role_kubernetes" {
  name = "irsa-role-kubernetes"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Principal = {
          Federated = "arn:aws:iam::${local.account_id}:oidc-provider/${local.oidc_provider_url}"
        },
        Action = "sts:AssumeRoleWithWebIdentity",
        Condition = {
          StringEquals = {
            "${local.oidc_provider_url}:sub" = "system:serviceaccount:default:${local.sa_name}"
          }
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "attach_irsa_s3_policy" {
  role       = aws_iam_role.irsa_role_kubernetes.name
  policy_arn = aws_iam_policy.gomu_landing_bucket_policy.arn
}
