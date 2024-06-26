resource "aws_iam_role" "ecs_task_execution_role" {
  name = "ecs_task_execution_role"

  assume_role_policy = <<EOF
{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Action": "sts:AssumeRole",
     "Principal": {
       "Service": "ecs-tasks.amazonaws.com"
     },
     "Effect": "Allow",
     "Sid": ""
   }
 ]
}
EOF
}

resource "aws_iam_role" "ecs_task_role" {
  name = "ecs_task_role"

  assume_role_policy = <<EOF
{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Action": "sts:AssumeRole",
     "Principal": {
       "Service": "ecs-tasks.amazonaws.com"
     },
     "Effect": "Allow",
     "Sid": ""
   }
 ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "ecs-task-execution-role-policy-attachment" {
  role       = aws_iam_role.ecs_task_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_ecs_cluster" "ecs_cluster" {
  name = "gomu-ecs-cluster"
}

resource "aws_ecs_task_definition" "extract_definition" {
  family                   = "gomu-task-definition"
  task_role_arn            = aws_iam_role.ecs_task_role.arn
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  network_mode             = "awsvpc"
  cpu                      = "256"
  memory                   = "1024"
  requires_compatibilities = ["FARGATE"]
  runtime_platform {
    cpu_architecture = "X86_64"
  }
  container_definitions = <<DEFINITION
[
  {
    "image": "${local.account_id}.dkr.ecr.${local.aws_region}.amazonaws.com/${aws_ecr_repository.gomu_repo.name}:latest",
    "name": "gomu-gomu",
    "logConfiguration": {
                "logDriver": "awslogs",
                "options": {
                    "awslogs-region" : "${local.aws_region}",
                    "awslogs-group" : "${aws_cloudwatch_log_group.gomu_log_group.name}",
                    "awslogs-stream-prefix" : "gomu"
                }
    },
    "secrets": [
            {
                "name": "CLIENT_ID",
                "valueFrom": "arn:aws:ssm:${local.aws_region}:${local.account_id}:parameter/client_id"
            },
            {
                "name": "CLIENT_SECRET",
                "valueFrom": "arn:aws:ssm:${local.aws_region}:${local.account_id}:parameter/client_secret"
            },
            {
                "name": "REFRESH_TOKEN",
                "valueFrom": "arn:aws:ssm:${local.aws_region}:${local.account_id}:parameter/refresh_token"
            }
    ],
    "environment": [
            {
                "name": "vpc_id",
                "value": "${local.vpc_id}"
            },
            {
                "name": "first_subnet_id",
                "value": "${local.first_subnet_id}"
            },
            {
                "name": "second_subnet_id",
                "value": "${local.second_subnet_id}"
            },
            {
                "name": "security_group_id",
                "value": "${local.security_group_id}"
            }
    ]
  }
]
DEFINITION
}

resource "aws_ecs_task_definition" "load_definition" {
  family                   = "sabo-task-definition"
  task_role_arn            = aws_iam_role.ecs_task_role.arn
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  network_mode             = "awsvpc"
  cpu                      = "256"
  memory                   = "1024"
  requires_compatibilities = ["FARGATE"]
  runtime_platform {
    cpu_architecture = "X86_64"
  }
  container_definitions = <<DEFINITION
[
  {
    "image": "${local.account_id}.dkr.ecr.${local.aws_region}.amazonaws.com/${aws_ecr_repository.sabo_repo.name}:latest",
    "name": "sabo-sabo",
    "logConfiguration": {
                "logDriver": "awslogs",
                "options": {
                    "awslogs-region" : "${local.aws_region}",
                    "awslogs-group" : "${aws_cloudwatch_log_group.gomu_log_group.name}",
                    "awslogs-stream-prefix" : "sabo"
                }
    },
    "secrets": [
            {
                "name": "SUPABASE_CONNECTION_ID",
                "valueFrom": "arn:aws:ssm:${local.aws_region}:${local.account_id}:parameter/supabase_connection_id"
            }
    ],
    "environment": [
            {
                "name": "vpc_id",
                "value": "${local.vpc_id}"
            },
            {
                "name": "first_subnet_id",
                "value": "${local.first_subnet_id}"
            },
            {
                "name": "second_subnet_id",
                "value": "${local.second_subnet_id}"
            },
            {
                "name": "security_group_id",
                "value": "${local.security_group_id}"
            }
    ]
  }
]
DEFINITION
}

resource "aws_ecs_task_definition" "transformation_definition" {
  family                   = "ace-task-definition"
  task_role_arn            = aws_iam_role.ecs_task_role.arn
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  network_mode             = "awsvpc"
  cpu                      = "256"
  memory                   = "1024"
  requires_compatibilities = ["FARGATE"]
  runtime_platform {
    cpu_architecture = "X86_64"
  }
  container_definitions = <<DEFINITION
[
  {
    "image": "${local.account_id}.dkr.ecr.${local.aws_region}.amazonaws.com/${aws_ecr_repository.ace_repo.name}:latest",
    "name": "ace-ace",
    "logConfiguration": {
                "logDriver": "awslogs",
                "options": {
                    "awslogs-region" : "${local.aws_region}",
                    "awslogs-group" : "${aws_cloudwatch_log_group.gomu_log_group.name}",
                    "awslogs-stream-prefix" : "ace"
                }
    },
    "secrets": [
            {
                "name": "DBT_HOST",
                "valueFrom": "arn:aws:ssm:${local.aws_region}:${local.account_id}:parameter/dbt_host"
            },
            {
                "name": "DBT_USER",
                "valueFrom": "arn:aws:ssm:${local.aws_region}:${local.account_id}:parameter/dbt_user"
            },
            {
                "name": "DBT_PASSWORD",
                "valueFrom": "arn:aws:ssm:${local.aws_region}:${local.account_id}:parameter/dbt_password"
            }
    ],
    "environment": [
            {
                "name": "vpc_id",
                "value": "${local.vpc_id}"
            },
            {
                "name": "first_subnet_id",
                "value": "${local.first_subnet_id}"
            },
            {
                "name": "second_subnet_id",
                "value": "${local.second_subnet_id}"
            },
            {
                "name": "security_group_id",
                "value": "${local.security_group_id}"
            }
    ]
  }
]
DEFINITION
}

resource "aws_iam_policy" "ssm_parameter_store_permissions" {
  name        = "parameter_read_permissions"
  description = "Allow "

  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "ParamsPolicy",
        "Effect" : "Allow",
        "Action" : [
          "ssm:GetParameters"
        ],
        "Resource" : [
          "arn:aws:ssm:${local.aws_region}:${local.account_id}:parameter/client_id",
          "arn:aws:ssm:${local.aws_region}:${local.account_id}:parameter/client_secret",
          "arn:aws:ssm:${local.aws_region}:${local.account_id}:parameter/refresh_token",
          "arn:aws:ssm:${local.aws_region}:${local.account_id}:parameter/supabase_connection_id",
          "arn:aws:ssm:${local.aws_region}:${local.account_id}:parameter/dbt_host",
          "arn:aws:ssm:${local.aws_region}:${local.account_id}:parameter/dbt_user",
          "arn:aws:ssm:${local.aws_region}:${local.account_id}:parameter/dbt_password"
        ]
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "attach_ssm_parameters_read_policy_to_ecs_task_execution_role" {
  role       = aws_iam_role.ecs_task_execution_role.name
  policy_arn = aws_iam_policy.ssm_parameter_store_permissions.arn
}
