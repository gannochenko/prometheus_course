run_infra: ## Run the infrastructure locally
	@docker-compose up --build

reload_config:
	@curl -X POST http://localhost:9090/-/reload
