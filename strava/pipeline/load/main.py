import argparse
from utils.pg_helper import SupaCursor
from utils.settings import S3_BUCKET_NAME


def main():

    parser = argparse.ArgumentParser(description="Load to Supabase")
    parser.add_argument(
        "schema_name", type=str, help="The schema the table should land in"
    )
    parser.add_argument("table_name", type=str, help="The name of the table")
    parser.add_argument("s3_key", type=str, help="The path of the S3 key")
    args = parser.parse_args()

    schema_name, table_name, s3_bucket, s3_key = (
        args.schema_name,
        args.table_name,
        S3_BUCKET_NAME,
        args.s3_key,
    )

    postgres_cursor = SupaCursor(schema_name, table_name, s3_bucket, s3_key)
    postgres_cursor.load()


if __name__ == "__main__":
    main()
