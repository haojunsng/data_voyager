import psycopg2
import boto3
import pandas as pd
import json
from utils.env_helper import load_env_vars
from sqlalchemy import create_engine


class SupaCursor:
    def __init__(self, table_name, s3_bucket, s3_key, if_exists="replace", index=False):
        self.table_name = table_name
        self.s3_bucket = s3_bucket
        self.s3_key = s3_key
        self.if_exists = if_exists
        self.index = index
        self.connection_string = load_env_vars()["SUPABASE_CONNECTION_ID"]
        self.conn = None
        self.cursor = None

    def load(self):

        try:
            s3 = boto3.client("s3")
            response = s3.get_object(Bucket=self.s3_bucket, Key=self.s3_key)
            data = json.loads(response["Body"].read().decode("utf-8"))
            df = pd.json_normalize(data)
        except Exception as e:
            print("Error retrieving data from S3:", e)
            return None

        try:
            db = create_engine(self.connection_string)
            df.to_sql(self.table_name, db, if_exists=self.if_exists, index=self.index)
            print("Loaded to Supabase!")
        except psycopg2.Error as e:
            print("Error: Unable to connect to the Supabase...")
            print(e)
