import os


# Administrative Information
OWNER = "monkeyDluffy"
EMAIL = "luffy@imu.com"

# ECS Configurations
GOMU_TASK_DEFINITION = "gomu-task-definition"
GOMU_CONTAINER_NAME = "gomu-gomu"
SABO_TASK_DEFINITION = "sabo-task-definition"
SABO_CONTAINER_NAME = "sabo-sabo"
ECS_CLUSTER = "gomu-ecs-cluster"
NETWORK_CONFIGURATION = {
    "awsvpcConfiguration": {
        "subnets": [
            os.environ.get("first_subnet_id"),
            os.environ.get("second_subnet_id"),
        ],
        "securityGroups": [os.environ.get("security_group_id")],
        "assignPublicIp": "ENABLED",
    }
}

# S3
S3_RAW_BUCKET = "gomu-landing-bucket"
