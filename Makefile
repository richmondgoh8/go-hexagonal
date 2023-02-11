.PHONY: postgres destroy

postgres:
	docker-compose -f docker/postgres/docker-compose.yml up -d -V

destroy:
	docker-compose --log-level ERROR -f docker/postgres/docker-compose.yml down

run:
	go run cmd/server.go