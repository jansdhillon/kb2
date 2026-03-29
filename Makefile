DB_GEN_PATH ?= $(shell pwd)/.gen
OAPI_DIR ?= $(shell pwd)/openapi

.PHONY: gen
gen:
	jet -dsn=postgresql://$(KB2_DB_USERNAME):$(KB2_DB_PASSWORD)@$(KB2_DB_HOST):$(KB2_DB_PORT)/$(KB2_DB_NAME)?sslmode=disable \
		-schema=$(KB2_DB_SCHEMA) \
		-path=$(GEN_PATH)

.PHONY: oapi-validate
oapi-validate:
	swagger-cli validate $(OAPI_DIR)/openapi.yaml

.PHONY: oapi-lint
oapi-lint:
	spectral lint $(OAPI_DIR)/openapi.yaml -F warn

.PHONY: oapi-bundle
oapi-bundle: oapi-validate
	swagger-cli bundle $(OAPI_DIR)/openapi.yaml -o $(OAPI_DIR)/openapi.bundle.yaml -t yaml

.PHONY: oapi-docs
oapi-docs:
	npx @redocly/cli build-docs $(OAPI_DIR)/landscape_api.bundle.yaml -o $(OAPI_DIR)/docs.html

.PHONY: oapi-serve-docs
oapi-serve-docs: oapi-docs
	python3 -m http.server 8080 --directory $(OAPI_DIR)

.PHONY: oapi-gen
oapi-gen: oapi-bundle
	go generate $(OAPI_DIR)/...
