from airflow import DAG
from utils.settings import OWNER, EMAIL
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
    "daily_at_0_utc",
    default_args=default_args,
    description="A DAG that runs daily at 00:00 UTC",
    schedule_interval="0 0 * * *",
    catchup=False,
)


S3_KEY = f"source=strava/type=activities/date={{{{ macros.ds_format(data_interval_end, '%Y-%m-%d %H:%M:%S%z', '%Y%m%d') }}}}/data.json"
# Define the task to print the current date and time
strava_to_s3 = StravaToS3Operator(
    task_id="strava_to_s3",
    dag=dag,
    s3_key=S3_KEY,
)

s3_stats_to_supabase = S3ToSupabaseOperator(
    task_id="load_strava_activity",
    dag=dag,
    table_name="ods_strava__activity",
    s3_key=S3_KEY,
)

strava_to_s3 >> s3_stats_to_supabase
