import requests
from utils.get_access_token import get_access_token


def get(endpoint):

    headers = {"Authorization": f"Authorization: Bearer {get_access_token()}"}
    try:
        response = requests.get(endpoint, headers=headers)
        response.raise_for_status()
    except requests.RequestException as e:
        print("Error making API call to Strava:", e)
        return None

    print("Done")
    return response.json()
