.PHONY: build_services launch_services stop_services

build_services:
	docker compose build

launch_services:
	docker compose up --force-recreate --abort-on-container-exit

stop_services:
	docker compose down -v
