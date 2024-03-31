import requests
from utils.get_access_token import get_access_token


ID = "47247266"
STATS_ENDPOINT = f"https://www.strava.com/api/v3/athletes/{ID}/stats"


def main() -> dict:
    headers = {"Authorization": f"Authorization: Bearer {get_access_token()}"}
    response = requests.get(STATS_ENDPOINT, headers=headers)
    response.raise_for_status()
    activity_data = response.json()
    return activity_data


if __name__ == "__main__":
    main()
