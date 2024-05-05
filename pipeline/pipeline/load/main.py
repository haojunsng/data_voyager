import argparse
from utils.pg_helper import SupaCursor


def main():

    parser = argparse.ArgumentParser(description="Load to Supabase")
    parser.add_argument("table_name", type=str, help="The name of the table")
    parser.add_argument("s3_bucket", type=str, help="The name of the S3 bucket")
    parser.add_argument("s3_key", type=str, help="The path of the S3 key")
    args = parser.parse_args()

    table_name, s3_bucket, s3_key = args.table_name, args.s3_bucket, args.s3_key

    postgres_cursor = SupaCursor(table_name, s3_bucket, s3_key)
    postgres_cursor.load()


if __name__ == "__main__":
    main()
