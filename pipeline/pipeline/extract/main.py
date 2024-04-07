from utils.s3_helper import write_to_s3
from utils.pull_from_strava import pull_stats_from_strava


ID = "47247266"


def main():
    data = pull_stats_from_strava(ID)
    write_to_s3(data)


if __name__ == "__main__":
    main()
