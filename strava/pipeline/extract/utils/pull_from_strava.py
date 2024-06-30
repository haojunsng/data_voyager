from utils.api_methods import get
import datetime


def pull_stats_from_strava(id):

    print(f"Pulling statistics...")
    STATS_ENDPOINT = f"https://www.strava.com/api/v3/athletes/{id}/stats"
    return get(STATS_ENDPOINT)


def pull_athlete_from_strava():

    print(f"Pulling athlete information...")
    ATHLETE_ENDPOINT = "https://www.strava.com/api/v3/athlete"
    return get(ATHLETE_ENDPOINT)


def pull_activity_from_strava(all=False, **kwargs):

    if not all:
        if "start" not in kwargs or "end" not in kwargs:
            raise ValueError("Missing 'start' and/or 'end' parameters")
        start = kwargs["before"]
        end = kwargs["after"]

        print(f"Pulling activities between {start} and {end}")
        ACTIVITY_ENDPOINT = f"https://www.strava.com/api/v3/athlete/activities?before={end}&after={start}&page=&per_page="

    else:
        now = int(datetime.datetime.now().timestamp())

        print("Pulling all activities")
        ACTIVITY_ENDPOINT = f"https://www.strava.com/api/v3/athlete/activities?before={now}&after=0&page=1&per_page=200"

    return get(ACTIVITY_ENDPOINT)
