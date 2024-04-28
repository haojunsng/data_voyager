# Makefile for local Airflow Setup
setup:
	pip install -r requirements.txt

airflow-init:
	docker compose up airflow-init

airflow-start:
	docker compose up

airflow-cleanup:
	docker compose down --volumes --remove-orphans
