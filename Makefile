.PHONY: pre-commit

pre-commit:
	@echo "Setting up pre-commit hooks..."
	pre-commit install --hook-type pre-commit --hook-type pre-push
