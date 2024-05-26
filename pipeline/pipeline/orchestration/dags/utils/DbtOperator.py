from airflow.providers.amazon.aws.operators.ecs import EcsRunTaskOperator
from .settings import (
    ACE_TASK_DEFINITION,
    ECS_CLUSTER,
    ACE_CONTAINER_NAME,
    NETWORK_CONFIGURATION,
)


class DbtOperator(EcsRunTaskOperator):
    """
    Custom DbtOperator that inherits from EcsRunTaskOperator
    """

    def __init__(self, cmd, select=None, **kwargs):
        super().__init__(
            task_definition=ACE_TASK_DEFINITION,
            cluster=ECS_CLUSTER,
            overrides={
                "containerOverrides": [
                    {
                        "name": ACE_CONTAINER_NAME,
                        "command": [f"dbt {cmd} --select {select}"],
                    },
                ],
            },
            network_configuration=NETWORK_CONFIGURATION,
            **kwargs,
        )

    def execute(self, context):
        super().execute(context)
