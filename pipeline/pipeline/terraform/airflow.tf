module "mwaa" {
  count   = var.deploy_mwaa ? 1 : 0
  source  = "aws-ia/mwaa/aws"
  version = "0.0.4"

  name              = "haoshoku"
  airflow_version   = "2.6.3"
  environment_class = "mw1.small"

  vpc_id             = var.vpc_id
  private_subnet_ids = [var.first_subnet_id, var.second_subnet_id]

  min_workers           = 1
  max_workers           = 5
  webserver_access_mode = "PRIVATE_ONLY"

  logging_configuration = {
    dag_processing_logs = {
      enabled   = true
      log_level = "INFO"
    }

    scheduler_logs = {
      enabled   = true
      log_level = "INFO"
    }

    task_logs = {
      enabled   = true
      log_level = "INFO"
    }

    webserver_logs = {
      enabled   = true
      log_level = "INFO"
    }

    worker_logs = {
      enabled   = true
      log_level = "INFO"
    }
  }

  airflow_configuration_options = {
    "core.load_default_connections" = "false"
    "core.load_examples"            = "false"
    "webserver.dag_default_view"    = "tree"
    "webserver.dag_orientation"     = "TB"
    "logging.logging_level"         = "INFO"
  }
}
