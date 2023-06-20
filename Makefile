# Setup applications dependencies:
setup-local:
	docker-compose --file docker-compose.yml up -d

# Run run:
run:
	go run cmd/main.go

# Run migrations:
migrate:
	migrate -path db/migrations -database ${DB_URL} up

# Run build:
build:
	go build -o bin/main cmd/main.go

# Run integrations tests:
integration-test:
	echo "Running integration tests"
	go test -tags="integration" -timeout=5m -v ./...