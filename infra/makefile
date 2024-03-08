DOCKER_COMPOSE=docker compose

help:
	@grep -E '(^[a-zA-Z0-9\./_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

run_infra: ## runs kafka
	$(DOCKER_COMPOSE) -f ./compose-kafka.yml up

stop_infra: ## stops kafka
	$(DOCKER_COMPOSE) -f ./compose-kafka.yml down


run_scrape: ## runs the scrapers 
	$(DOCKER_COMPOSE) -f ../infra/compose-scrape.yml up

stop_scrape: # stops the scraper services
	$(DOCKER_COMPOSE) -f ../infra/compose-scrape.yml up

.PHONY: help run_site run_scrape run_infra stop_infra stop_scrape

