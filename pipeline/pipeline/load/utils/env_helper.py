import os


def load_env_vars() -> dict:

    SUPABASE_CONNECTION_ID = os.environ.get("SUPABASE_CONNECTION_ID")

    env_vars = {
        "SUPABASE_CONNECTION_ID": SUPABASE_CONNECTION_ID,
    }

    return env_vars
