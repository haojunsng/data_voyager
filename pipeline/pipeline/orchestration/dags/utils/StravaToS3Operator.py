from airflow.providers.amazon.aws.operators.ecs import EcsRunTaskOperator
import os
from settings import (
    TASK_DEFINITION,
    ECS_CLUSTER,
    ECS_CONTAINER_NAME,
    AWSLOGS_GROUP,
    REGION,
)


class StravaToS3Operator(EcsRunTaskOperator):
    """
    Custom StravaToS3Operator that inherits from EcsRunTaskOperator
    """

    # Network Configuration has been passed into the ECS Container
    subnets = [os.environ.get("first_subnet_id"), os.environ.get("second_subnet_id")]
    securityGroups = [os.environ.get("security_group_id")]
    NETWORK_CONFIGURATION = {
        "awsvpcConfiguration": {
            "assignPublicIp": "ENABLED",
            "securityGroups": securityGroups,
            "subnets": subnets,
        }
    }

    def __init__(self, strava_id, **kwargs):
        super().__init__(
            task_definition=TASK_DEFINITION,
            cluster=ECS_CLUSTER,
            overrides={
                "containerOverrides": [
                    {
                        "name": ECS_CONTAINER_NAME,
                        "command": [f"python main.py {strava_id}"],
                    },
                ],
            },
            network_configuration=self.NETWORK_CONFIGURATION,
            awslogs_group=AWSLOGS_GROUP,
            awslogs_region=REGION,
            **kwargs,
        )

    def execute(self, context):
        super().execute(context)
