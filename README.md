# strava_pipeline

## Initial Data Architecture Diagram
![Archi](https://github.com/haojunsng/simple_pipeline/blob/main/pipeline/assets/archi.png)

## Repository Navigation
This repository contains 5 parts - navigate to the respective subfolders for a more in-depth description.

I chose to adopt a monorepo approach only because this is more of an exploratory/hobby work and did not want the hassle of maintaining multiple repositories.


- [`extract/`](https://github.com/haojunsng/strava_pipeline/tree/main/pipeline/pipeline/extract) contains the logic of data extraction from STRAVA API.
- [`load/`](https://github.com/haojunsng/strava_pipeline/tree/main/pipeline/pipeline/load) contains the logic of loading data from landing buckets to database.
- [`transformation/`](https://github.com/haojunsng/strava_pipeline/tree/main/pipeline/pipeline/transformation) contains the transformation logic.
- [`orchestration/`](https://github.com/haojunsng/strava_pipeline/tree/main/pipeline/pipeline/orchestration) contains the airflow code and DAGs.
- [`iac/`](https://github.com/haojunsng/strava_pipeline/tree/main/pipeline/pipeline/iac) contains the IaC for all necessary resources provisioned.

### `extract`
---
#### Description
1. Obtain the following credentials from STRAVA App Integration:
    - `CLIENT_ID`
    - `CLIENT_SECRET`
    - `REFRESH_TOKEN`
2. Store credentials in AWS SSM Parameter Store.

3. Implement logic of data extraction in `main.py`.

4. Pass credentials into container through ECS Task Definition.

5. Dockerise `extract/` and push to ECR. CD through GHA has been implemented in `cd.yaml` to update image upon merging to `main`.

6. Compute will be orchestrated by Airflow through custom operator `StravaToS3Operator` which inherits from `ECSRunTaskOperator`.

### `load`
---
#### Description
[Supabase](https://supabase.com/) is chosen as the postgres database for this project mostly because they recently went GA, and the UI looks pretty clean and most importantly I can keep within the free tier very comfortably.

#### Supabase
![supabase](https://github.com/haojunsng/simple_pipeline/blob/main/pipeline/assets/supabase.png)
- Data loaded from S3.

### `transformation`
---
[dbt](https://docs.getdbt.com/docs/introduction) is chosen to handle all data transformation work required.

### `orchestration`
---
#### Description
[Airflow](https://airflow.apache.org/) is chosen to manage all orchestration work around extracting, loading and transforming of data.
- Extraction of data from STRAVA API to S3 bucket: [StravaToS3Operator](https://github.com/haojunsng/strava_pipeline/blob/main/pipeline/pipeline/orchestration/dags/utils/StravaToS3Operator.py)
- Loading of data from S3 bucket to Supabase: [S3ToSupabaseOperator](https://github.com/haojunsng/strava_pipeline/blob/main/pipeline/pipeline/orchestration/dags/utils/S3ToSupabaseOperator.py)

#### `StravaToS3Operator` & `S3ToSupabaseOperator`:
---
![dag](https://github.com/haojunsng/simple_pipeline/blob/main/pipeline/assets/dag.png)
- Custom [StravaToS3Operator](https://github.com/haojunsng/strava_pipeline/blob/main/pipeline/pipeline/orchestration/dags/utils/StravaToS3Operator.py) inherits EcsRunTaskOperator and is created to call the STRAVA API for extraction.
- Similarly, custom [S3ToSupabaseOperator](https://github.com/haojunsng/strava_pipeline/blob/main/pipeline/pipeline/orchestration/dags/utils/S3ToSupabaseOperator.py) also inherits EcsRunTaskOperator and helps to load data from my S3 bucket to Supabase Postgres database.
- Both logic (STRAVA extraction & Loading to Supabase) are managed in `extract/` and `load/` respectively

#### Deployment of DAGs to Airflow:
---
![s3](https://github.com/haojunsng/simple_pipeline/blob/main/pipeline/assets/s3.png)
- Deployed to AWS S3 bucket through Github Actions `aws s3 sync` for MWAA cluster


#### `dev/`:
---
- Local airflow development environment for testing
- Symlinked to `'orchestration/dags/`
- Use `docker-compose up` to spin up local airflow

### `iac`
---
#### Description
[Terraform](https://www.terraform.io/) is chosen to support the IaC for this entire strava pipeline project.

#### Resources maintained using Terraform:
- Networking
    - VPC
    - Subnets
    - Security Group
- ECS Task Definition
- ECR
- Identity Access Management
    - Service User for GHA
    - ECS Task Execution Role
    - ECS Task Role
    - Respective IAM policies required around authorisation management
- Cloudwatch Logs
- S3 Buckets

#### Resources NOT maintained using Terraform:
- SSM Parameters

#### Caveats
- The supabase terraform provider is very very new. Might not be as stable unfortunately.
- Currently only supports local `terraform init/plan/apply` as `.tfstate` files are maintained locally -- does not support collaboration yet.


### Environment Variables Management

#### Using direnv and .envrc

Selected [direnv](https://direnv.net/) along with a `.envrc` file to manage environment variables in the development environment. This automatically loads environment variables when entering the project directory.

#### Terraform Environment Variables

For Terraform-related environment variables, they are prefixed with `TF_VAR_`. This ensures that the environment variables can be registered by `.tf` files.

#### Python Environment Variables

In Python code, a straightforward method of retrieval is used (`os.environ.get`) to access environment variables.

### Secure Credential Management

Confidential credentials, such as API keys and database passwords, are securely stored in the AWS Systems Manager (SSM) Parameter Store. They are retrieved securely at runtime and passed into the ECS containers.
