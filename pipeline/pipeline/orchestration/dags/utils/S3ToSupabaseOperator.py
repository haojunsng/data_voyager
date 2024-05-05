from airflow.providers.amazon.aws.operators.ecs import EcsRunTaskOperator
from .settings import (
    SABO_TASK_DEFINITION,
    ECS_CLUSTER,
    SABO_CONTAINER_NAME,
    NETWORK_CONFIGURATION,
)


class S3ToSupabaseOperator(EcsRunTaskOperator):
    """
    Custom S3ToSupabaseOperator that inherits from EcsRunTaskOperator
    """

    def __init__(self, **kwargs):
        super().__init__(
            task_definition=SABO_TASK_DEFINITION,
            cluster=ECS_CLUSTER,
            overrides={
                "containerOverrides": [
                    {
                        "name": SABO_CONTAINER_NAME,
                        "command": ["python main.py"],
                    },
                ],
            },
            network_configuration=NETWORK_CONFIGURATION,
            **kwargs,
        )

    def execute(self, context):
        super().execute(context)
