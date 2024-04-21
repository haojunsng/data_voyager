from utils.s3_helper import write_to_s3
from utils.pull_from_strava import (
    pull_stats_from_strava,
    pull_activity_from_strava,
    pull_athlete_from_strava,
)


ID = "47247266"


def main():

    # Pull snapshot of statistics of athlete - partitioned by `date`
    stats = pull_stats_from_strava(ID)
    write_to_s3(stats, "statistics")

    # Pull snapshot of all activities of athlete - partitioned by `date`
    activities = pull_activity_from_strava(all=True)
    write_to_s3(activities, "activities")


if __name__ == "__main__":
    main()
