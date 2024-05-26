select
    activity_id,
    name,
    sport_type,
    distance,
    moving_time,
    elapsed_time,
    total_elevation_gain,
    average_heartrate,
    max_heartrate,
    suffer_score,
    average_temp
from
    {{ ref('dwd_strava__activity_df') }}
