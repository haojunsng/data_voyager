import argparse
from utils.s3_helper import write_to_s3
from utils.pull_from_strava import (
    pull_activity_from_strava,
)


def main():

    parser = argparse.ArgumentParser(description="Extract STRAVA")
    parser.add_argument("S3_KEY", type=str, help="The S3_KEY to land in")
    args = parser.parse_args()
    S3_KEY = args.S3_KEY

    # Pull snapshot of all activities of athlete - partitioned by `date`
    activities = pull_activity_from_strava(all=True)
    write_to_s3(activities, S3_KEY)


if __name__ == "__main__":
    main()
