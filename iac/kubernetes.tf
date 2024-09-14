resource "kubernetes_deployment" "weather_app" {
  metadata {
    name      = "weather-app"
    namespace = "default"
  }

  spec {
    replicas = 2

    selector {
      match_labels = {
        app = "weather-app"
      }
    }

    template {
      metadata {
        labels = {
          app = "weather-app"
        }
      }

      spec {
        container {
          name    = "weather-producer"
          image   = "${aws_ecr_repository.koby_repo.repository_url}:latest"
          command = ["go", "run", "./producer"]

          port {
            container_port = 8081
          }

          env {
            name = "KAFKA_BROKER"
            value_from {
              secret_key_ref {
                name = kubernetes_secret.kafka_broker.metadata[0].name
                key  = "kafka_broker"
              }
            }
          }
        }

        container {
          name    = "weather-consumer"
          image   = "${aws_ecr_repository.koby_repo.repository_url}:latest"
          command = ["go", "run", "./consumer"]

          port {
            container_port = 8082
          }

          env {
            name = "KAFKA_BROKER"
            value_from {
              secret_key_ref {
                name = kubernetes_secret.kafka_broker.metadata[0].name
                key  = "kafka_broker"
              }
            }
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "weather_service" {
  metadata {
    name      = "weather-service"
    namespace = "default"
  }

  spec {
    selector = {
      app = "weather-app"
    }

    port {
      protocol    = "TCP"
      port        = 8081
      target_port = 8081
    }

    port {
      protocol    = "TCP"
      port        = 8082
      target_port = 8082
    }
  }
}
