.PHONY: pre-commit

setup:
	@echo "Installing necessary packages for this repository..."
	brew install $(<packages.txt)

pre-commit:
	@echo "Setting up pre-commit hooks..."
	pre-commit install --hook-type pre-commit --hook-type pre-push

test:
	@echo "Running tests..."
	cd orchestration/dags && python tests/dag_integrity.py
