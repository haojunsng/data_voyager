.PHONY: pre-commit

setup:
	@echo "Installing necessary packages for this repository..."
	brew bundle

pre-commit:
	@echo "Setting up pre-commit hooks..."
	pre-commit install --hook-type pre-commit --hook-type pre-push

test:
	@echo "Running tests..."
	cd strava/pipeline/orchestration/dags && python tests/dag_integrity.py
