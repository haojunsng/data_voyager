import requests
from utils.env_helper import load_env_vars


AUTH_ENDPOINT = "https://www.strava.com/oauth/token"


def get_access_token() -> str:
    payload = {
        "client_id": load_env_vars()["CLIENT_ID"],
        "client_secret": load_env_vars()["CLIENT_SECRET"],
        "refresh_token": load_env_vars()["REFRESH_TOKEN"],
        "grant_type": "refresh_token",
        "f": "json",
    }
    res = requests.post(AUTH_ENDPOINT, data=payload, verify=True)
    access_token = res.json()["access_token"]
    return access_token
