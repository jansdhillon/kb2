GEN_PATH ?= $(shell pwd)/.gen

.PHONY: gen

gen:
	jet -dsn=postgresql://$(KB2_DB_USERNAME):$(KB2_DB_PASSWORD)@$(KB2_DB_HOST):$(KB2_DB_PORT)/$(KB2_DB_NAME)?sslmode=disable \
		-schema=$(KB2_DB_SCHEMA) \
		-path=$(GEN_PATH)

.PHONY: validate
validate:
	swagger-cli validate openapi/openapi.yaml

.PHONY: lint
lint:
	spectral lint openapi/openapi.yaml -F warn

.PHONY: bundle
bundle:
	swagger-cli bundle openapi/openapi.yaml -o openapi/landscape_api.bundle.yaml -t yaml

.PHONY: docs
docs:
	npx @redocly/cli build-docs openapi/landscape_api.bundle.yaml -o openapi/docs.html

.PHONY: serve-docs
serve-docs: docs
	python3 -m http.server 8080 --directory openapi
