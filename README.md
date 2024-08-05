# data_voyager

## Data Architecture
![image](https://github.com/haojunsng/data_voyager/assets/51106107/780bf52b-b46f-49f7-9a53-46c668e7694f)


## Repository Navigation
This repository contains 2 main parts - `strava/` and `weather/`

I chose to adopt a monorepo approach only because this is more of an exploratory/hobby work and did not want the hassle of maintaining multiple repositories.


[`strava/`](#strava) contains all code around the strava pipeline.
- A batch ELT data pipeline in Python, connecting to Postgres DB, orchestrated by Airflow and dbt (through ECS).

[`weather/`](#weather) contains all code around the weather pipeline.
- A near-realtime data pipeline in Golang utilizing Kafka on a Kubernetes Service, connecting to MongoDB, with Terraform as the IaC.

[`iac/`](#iac) contains all Terraform (chosen IaC) code.
- All cloud resources with the exception of SSM Parameters are provisioned using Terraform.

---
<a name="strava"></a>
### `strava/`


- [`extract/`](https://github.com/haojunsng/data_voyager/tree/main/strava/pipeline/extract) contains the logic of data extraction from STRAVA API.
- [`load/`](https://github.com/haojunsng/data_voyager/tree/main/strava/pipeline/load) contains the logic of loading data from landing buckets to database.
- [`transformation/`](https://github.com/haojunsng/data_voyager/tree/main/strava/pipeline/transformation) contains the transformation logic.
- [`orchestration/`](https://github.com/haojunsng/data_voyager/tree/main/strava/pipeline/orchestration) contains the airflow code and DAGs.


#### `extract`
---
##### Description
1. Obtain the following credentials from STRAVA App Integration:
    - `CLIENT_ID`
    - `CLIENT_SECRET`
    - `REFRESH_TOKEN`
2. Store credentials in AWS SSM Parameter Store.

3. Implement logic of data extraction in `main.py`.

4. Pass credentials into container through ECS Task Definition.

5. Dockerise `extract/` and push to ECR. CD through GHA has been implemented in `cd.yaml` to update image upon merging to `main`.

6. Compute will be orchestrated by Airflow through custom operator `StravaToS3Operator` which inherits from `ECSRunTaskOperator`.

#### `load`
---
##### Description
[Supabase](https://supabase.com/) is chosen as the postgres database for this project mostly because they recently went GA, and the UI looks pretty clean and most importantly I can keep within the free tier very comfortably.

##### Supabase
![supabase](https://github.com/haojunsng/data_voyager/blob/main/strava/assets/supabase.png)
- Data loaded from S3.

#### `transformation`
---
![dbt](https://github.com/haojunsng/data_voyager/blob/main/strava/assets/dbt.png)

[dbt](https://docs.getdbt.com/docs/introduction) is chosen to handle all data transformation work required.

##### dbt Project Management

A monorepo approach to dbt Project management is taken because there will be dependencies between `strava` dbt_project and `weather` dbt_project -- I'd prefer to have them all in 1 place just so the dependencies between can be captured by dbt.

#### `orchestration`
---
##### Description
[Airflow](https://airflow.apache.org/) is chosen to manage all orchestration work around extracting, loading and transforming of data.
- Extraction of data from STRAVA API to S3 bucket: [StravaToS3Operator](https://github.com/haojunsng/data_voyager/blob/main/strava/pipeline/orchestration/dags/utils/StravaToS3Operator.py)
- Loading of data from S3 bucket to Supabase: [S3ToSupabaseOperator](https://github.com/haojunsng/data_voyager/blob/main/strava/pipeline/orchestration/dags/utils/S3ToSupabaseOperator.py)

##### `StravaToS3Operator` & `S3ToSupabaseOperator`:
![dag](https://github.com/haojunsng/data_voyager/blob/main/strava/assets/dag.png)
- Custom [StravaToS3Operator](https://github.com/haojunsng/data_voyager/blob/main/strava/pipeline/orchestration/dags/utils/StravaToS3Operator.py) inherits EcsRunTaskOperator and is created to call the STRAVA API for extraction.
- Similarly, custom [S3ToSupabaseOperator](https://github.com/haojunsng/data_voyager/blob/main/strava/pipeline/orchestration/dags/utils/S3ToSupabaseOperator.py) also inherits EcsRunTaskOperator and helps to load data from my S3 bucket to Supabase Postgres database.
- Lastly, [DbtOperator](https://github.com/haojunsng/data_voyager/blob/main/strava/pipeline/orchestration/dags/utils/DbtOperator.py) which triggers dbt tasks through ECS to execute the transformation logic.
- All 3 logic (STRAVA extraction, Loading to Supabase & dbt Transformation) are managed in `extract/`, `load/` and `transformation/` respectively.

##### Deployment of DAGs to Airflow:
![s3](https://github.com/haojunsng/data_voyager/blob/main/strava/assets/s3.png)
- Deployed to AWS S3 bucket through Github Actions `aws s3 sync` for MWAA cluster


##### `dev/`:
- Local airflow development environment for testing
- Symlinked to `orchestration/dags/`
- Use `docker-compose up` to spin up local airflow

---
<a name="weather"></a>
### `weather/`

#### Kafka Producer
This was implemented in golang with [Open-Meteo API](https://github.com/innotechdevops/openmeteo).

![comment](https://github.com/haojunsng/data_voyager/blob/main/weather/assets/kafka_ui_live_messages.png)

---

<a name="iac"></a>
### `iac/`
##### Description
[Terraform](https://www.terraform.io/) is chosen to support the IaC for this entire strava pipeline project.

#### Resources maintained using Terraform:
- ECS Task Definition
- ECR
- Cloudwatch Logs
- S3 Buckets
- Networking
    - VPC
    - Subnets
    - Security Group
- Identity Access Management
    - Service User for GHA
    - ECS Task Execution Role
    - ECS Task Role
    - Respective IAM policies required around authorisation management

#### Resources NOT maintained using Terraform:
- SSM Parameters

##### Scalr

![comment](https://github.com/haojunsng/data_voyager/blob/main/strava/assets/scalr_ui.png)

Scalr was chosen to support remote terraform operations. The free tier supports up to <u>50 terraform operations monthly</u>.

`terraform plan` will execute upon raising a PR with commits from the declared directory -- `iac/`.

![comment](https://github.com/haojunsng/data_voyager/blob/main/strava/assets/scalr_comment.png)

`auto apply` has been disabled and plans have to be manually approved on the Scalr UI, which can be navigated from the PR comments.

![comment](https://github.com/haojunsng/data_voyager/blob/main/strava/assets/scalr_ci.png)

### Using GitHub Workflows with OIDC to Push Images to Amazon ECR

In this project, GitHub Workflows along with OIDC authentication are leveraged to automate the process of pushing/updating images to ECR on AWS.

### Environment Variables Management

#### Using direnv and .envrc

Selected [direnv](https://direnv.net/) along with a `.envrc` file to manage environment variables in the development environment. This automatically loads environment variables when entering the project directory.

#### Terraform Environment Variables

For Terraform-related environment variables, they are prefixed with `TF_VAR_`. This ensures that the environment variables can be registered by `.tf` files.

#### Python Environment Variables

In Python code, a straightforward method of retrieval is used (`os.environ.get`) to access environment variables.

### Secure Credential Management

Confidential credentials, such as API keys and database passwords, are securely stored in the AWS Systems Manager (SSM) Parameter Store. They are retrieved securely at runtime and passed into the ECS containers.

### Workflow Management
[Jira](https://snghaojun18.atlassian.net/jira/software/projects/SNG/boards/2)
<img width="966" alt="image" src="https://github.com/haojunsng/data_voyager/assets/51106107/e6e24f95-b634-49f9-8298-e6512418eab9">
