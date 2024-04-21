import boto3
import json
import datetime
from utils.settings import S3_BUCKET_NAME


def write_to_s3(data, source_type):

    try:
        s3 = boto3.client("s3")
        S3_KEY = f'source=strava/type={source_type}/date={datetime.datetime.now().strftime("%Y%m%d")}/stats.json'  # TODO: Replace with Airflow ds
        s3.put_object(
            Body=json.dumps(data, indent=4), Bucket=S3_BUCKET_NAME, Key=S3_KEY
        )
    except Exception as e:
        print("Error writing data to S3:", e)
        return None
