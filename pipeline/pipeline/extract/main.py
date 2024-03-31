import requests

from utils.get_access_token import get_access_token
from utils.s3_helper import write_to_s3


ID = "47247266"
STATS_ENDPOINT = f"https://www.strava.com/api/v3/athletes/{ID}/stats"


def main():

    # Pull from STRAVA
    headers = {"Authorization": f"Authorization: Bearer {get_access_token()}"}
    response = requests.get(STATS_ENDPOINT, headers=headers)
    response.raise_for_status()
    data = response.json()

    write_to_s3(data)


if __name__ == "__main__":
    main()
