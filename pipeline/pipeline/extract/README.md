# `extract/`

## Description
1. Obtain the following credentials from STRAVA App Integration:
    - `CLIENT_ID`
    - `CLIENT_SECRET`
    - `REFRESH_TOKEN`
2. Store credentials in AWS SSM Parameter Store.

3. Implement logic of data extraction in `main.py`.

4. Pass credentials into container through ECS Task Definition.

5. Dockerise `extract/` and push to ECR. CD through GHA has been implemented in `cd.yaml` to update image upon merging to `main`.

6. Compute will be orchestrated by Airflow through `ECSRunTaskOperator`.
