from airflow import DAG
from airflow.operators.dummy_operator import DummyOperator
from utils.settings import OWNER, EMAIL
from datetime import datetime, timedelta


default_args = {
    "owner": OWNER,
    "depends_on_past": False,
    "start_date": datetime(2024, 4, 30, 0, 0),
    "email": EMAIL,
    "email_on_failure": False,
    "email_on_retry": False,
    "retries": 1,
    "retry_delay": timedelta(minutes=5),
}

# Define the DAG object
dag = DAG(
    "daily_at_0000_utc",
    default_args=default_args,
    description="A DAG that runs daily at 00:00 UTC",
    schedule_interval="0 0 * * *",
    catchup=False,
)


# Define the task to print the current date and time
usopp = DummyOperator(
    task_id="dummy",
    dag=dag,
)
