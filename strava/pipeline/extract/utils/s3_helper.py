import boto3
import json
from utils.settings import S3_BUCKET_NAME


def write_to_s3(data, S3_KEY):

    try:
        s3 = boto3.client("s3")
        s3.put_object(
            Body=json.dumps(data, indent=4), Bucket=S3_BUCKET_NAME, Key=S3_KEY
        )
    except Exception as e:
        print("Error writing data to S3:", e)
        return None
