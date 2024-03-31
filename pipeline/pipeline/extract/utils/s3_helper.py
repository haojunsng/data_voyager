import boto3
import json
import datetime


def write_to_s3(data):
    s3 = boto3.client("s3")
    S3_BUCKET_NAME = "gomu-landing-bucket"
    S3_KEY = f'source=strava/type=statistics/date={datetime.datetime.now().strftime("%Y%m%d")}/stats.json'  # TODO: Replace with Airflow ds
    s3.put_object(Body=json.dumps(data, indent=4), Bucket=S3_BUCKET_NAME, Key=S3_KEY)
