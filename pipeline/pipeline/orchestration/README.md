# `orchestration/`

## Description

### `dags/`:
---
![dag_example](https://github.com/haojunsng/simple_pipeline/blob/main/pipeline/assets/dag_example.png)
- Managed through Airflow
- Custom [StravaToS3Operator](https://github.com/haojunsng/strava_pipeline/blob/5a9b8ab31742d75aad87b7f7b7178d9c0ea04f41/pipeline/pipeline/orchestration/dags/utils/StravaToS3Operator.py) is created to manage STRAVA API extraction using Airflow Operator, by inheriting EcsRunTaskOperator.
![s3](https://github.com/haojunsng/simple_pipeline/blob/main/pipeline/assets/s3.png)
- Deployed to AWS S3 bucket through Github Actions `aws s3 sync` for MWAA cluster


### `dev/`:
---
- Local airflow development environment for testing
