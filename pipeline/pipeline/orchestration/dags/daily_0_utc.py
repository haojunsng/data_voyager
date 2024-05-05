from airflow import DAG
from utils.settings import OWNER, EMAIL, S3_RAW_BUCKET
from utils.StravaToS3Operator import StravaToS3Operator
from utils.S3ToSupabaseOperator import S3ToSupabaseOperator
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
strava_to_s3 = StravaToS3Operator(task_id="strava_to_s3", dag=dag, strava_id="47247266")
s3_stats_to_supabase = S3ToSupabaseOperator(
    task_id="s3_stats_to_supabase",
    dag=dag,
    table_name="data",
    s3_bucket=S3_RAW_BUCKET,
    s3_key="source=strava/type=statistics/date=20240505/stats.json",
)  # TODO: Use airflow data_interval template variables

strava_to_s3 >> s3_stats_to_supabase
