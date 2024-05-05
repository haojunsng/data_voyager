# strava_pipeline

## Initial Data Architecture Diagram
![Archi](https://github.com/haojunsng/simple_pipeline/blob/main/pipeline/assets/archi.png)

## Repository Navigation
This repository contains three parts - navigate to the respective subfolders for a more in-depth description.
- [`extract/`](https://github.com/haojunsng/strava_pipeline/tree/main/pipeline/pipeline/extract#readme) contains the logic of data extraction from STRAVA API.
- [`load/`](https://github.com/haojunsng/strava_pipeline/tree/main/pipeline/pipeline/load#readme) contains the logic of loading data from landing buckets to database.
- [`transformation/`](https://github.com/haojunsng/strava_pipeline/tree/main/pipeline/pipeline/transformation#readme) contains the transformation logic.
- [`orchestration/`](https://github.com/haojunsng/strava_pipeline/tree/main/pipeline/pipeline/orchestration#readme) contains the airflow code and DAGs.
- [`terraform/`](https://github.com/haojunsng/strava_pipeline/tree/main/pipeline/pipeline/iac#readme) contains the IaC for all necessary resources provisioned.


### Environment Variables Management

#### Using direnv and .envrc

Selected [direnv](https://direnv.net/) along with a `.envrc` file to manage environment variables in the development environment. This automatically loads environment variables when entering the project directory.

#### Terraform Environment Variables

For Terraform-related environment variables, they are prefixed with `TF_VAR_`. This ensures that the environment variables can be registered by `.tf` files.

#### Python Environment Variables

In Python code, a straightforward method of retrieval is used (`os.environ.get`) to access environment variables.

### Secure Credential Management

Confidential credentials, such as API keys and database passwords, are securely stored in the AWS Systems Manager (SSM) Parameter Store. They are retrieved securely at runtime and passed into the ECS containers.
