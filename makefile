DOCKER_COMPOSE=docker compose

help:
	@echo "invalid subcommand valid commands:"
	@echo "help run_scrape run_site"

run_infra:
	$(DOCKER_COMPOSE)

run_scrape: 
	$(DOCKER_COMPOSE) -f ./infra/compose-scrape.yml up

run_site: 
	$(DOCKER_COMPOSE)

.PHONY: help run_site run_scrape run_infra

