.PHONY: gen

DB_USERNAME ?= postgres
DB_PASSWORD ?= postgres
DB_HOST ?= localhost
PORT ?= 5432
DATABASE ?= postgres

gen:
	jet -dsn=postgresql://$(USERNAME):$(PASSWORD)@$(HOST):$(PORT)/$(DATABASE)?sslmode=disable -schema=public -path=./.gen
