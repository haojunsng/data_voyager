import requests
from utils.get_access_token import get_access_token


def pull_stats_from_strava(id):

    STATS_ENDPOINT = f"https://www.strava.com/api/v3/athletes/{id}/stats"

    headers = {"Authorization": f"Authorization: Bearer {get_access_token()}"}
    response = requests.get(STATS_ENDPOINT, headers=headers)
    response.raise_for_status()
    return response.json()
