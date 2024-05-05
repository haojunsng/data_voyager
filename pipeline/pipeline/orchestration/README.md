# `orchestration/`

## Description

### `dags/`:
---
![dag_example](https://github.com/haojunsng/simple_pipeline/blob/main/pipeline/assets/dag_example.png)
- Managed through Airflow
- Deployed to AWS S3 bucket through Github Actions `aws s3 sync` for MWAA cluster


### `dev/`:
---
- Local airflow development environment for testing
