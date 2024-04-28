# `orchestration/`

## Description

### `dags/`:
---
- Managed through Airflow
- Deployed to AWS S3 bucket through Github Actions `aws s3 sync` for MWAA cluster

### `dev/`:
---
- Local airflow development environment for testing
