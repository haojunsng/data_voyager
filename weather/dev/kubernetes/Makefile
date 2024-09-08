# Variables
IMAGE_NAME = weather:latest
CLUSTER_NAME = kind-cluster
DEPLOYMENT_FILE = deployment.yaml
SERVICE_FILE = service.yaml

build:
	docker build -t $(IMAGE_NAME) .

start-cluster:
	kind create cluster --name $(CLUSTER_NAME)

load-image:
	kind load docker-image $(IMAGE_NAME) --name $(CLUSTER_NAME)

# Apply Kubernetes manifests
deploy:
	kubectl apply -f $(DEPLOYMENT_FILE)
	kubectl apply -f $(SERVICE_FILE)

# Delete deployments and services
delete-resources:
	kubectl delete -f $(DEPLOYMENT_FILE) || true
	kubectl delete -f $(SERVICE_FILE) || true

# Delete KinD cluster
delete-cluster:
	kind delete cluster --name $(CLUSTER_NAME)

# Build, start cluster, load image, and deploy all in one
up: build start-cluster load-image deploy

down: delete-resources delete-cluster

.PHONY: build start-cluster load-image deploy delete-resources delete-cluster up down
