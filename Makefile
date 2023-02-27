.PHONY: full destroy run gen

full:
	docker-compose -f docker-compose.yml up -d -V --build

postgres:
	docker-compose -f docker-compose-postgres.yml up -d -V

destroy:
	docker-compose --log-level ERROR -f docker-compose.yml down --remove-orphans

run: swagger
	go run cmd/server.go

gen:
	mockgen -source=./internal/core/ports/ports.go -destination=./internal/mocks/core/ports/ports.go

swagger:
	which swagger || go install github.com/go-swagger/go-swagger
	swagger generate spec -o ./swagger.yaml --scan-models