# `orchestration/`

## Description
This subdirectory manages all orchestration work around extracting & loading.
- Extracting of data from STRAVA API to S3 bucket: [StravaToS3Operator](https://github.com/haojunsng/strava_pipeline/blob/5a9b8ab31742d75aad87b7f7b7178d9c0ea04f41/pipeline/pipeline/orchestration/dags/utils/StravaToS3Operator.py)
- Loading of data from S3 bucket to Supabase: [S3ToSupabaseOperator]

### `StravaToS3Operator` & `S3ToSupabaseOperator`:
---
![dag](https://github.com/haojunsng/simple_pipeline/blob/main/pipeline/assets/dag.png)
- Custom StravaToS3Operator is created to manage STRAVA API extraction using Airflow Operator, by inheriting EcsRunTaskOperator
- Custom S3ToSupabaseOperator is created to load data from S3 to Supabase Postgres database.
- Both logics (STRAVA extraction & Loading to Supabase are managed in `extract/` and `load/` respectively)

### Deployment of DAGs to Airflow:
---
![s3](https://github.com/haojunsng/simple_pipeline/blob/main/pipeline/assets/s3.png)
- Deployed to AWS S3 bucket through Github Actions `aws s3 sync` for MWAA cluster


### `dev/`:
---
- Local airflow development environment for testing
- Symlinked to `'orchestration/dags/`
- Use `docker-compose up` to spin up local airflow
