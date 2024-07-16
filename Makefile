THIS_FILE := $(lastword $(MAKEFILE_LIST))

.PHONY:  build up start down destroy stop logs-app

build:
	docker compose  build $(c)

up:
	docker compose  up -d $(c)

start:
	docker compose  start $(c)

down:
	docker compose  down $(c)

destroy:
	docker compose  down -v $(c)

stop:
	docker compose  stop $(c)

logs-app:
	docker compose  logs --tail=100 -f app