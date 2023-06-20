<h1 align="center">Airfright App</h1>

<div align="center">
  An app that is responsible for managing and transporting shipments to Warhouse Management System.
</div>

### Table of Contents
- [Structure](#structure)
- [How To Run](#how_to_run)
- [Integration Tests](#integration_test)

### Structure <a name = "structure"></a>
- `cmd` - Main folder contains `main` run application
- `config` - Contains configuration files
- `db` - Contains all the database migration
- `internal` - Contains all the business logic
- `docker-compose.yml` - Needed to run the application locally
- `Dockerfile` - Needed to build the application's image
- `go.mod` - Contains all the dependencies
- `Makefile` - Contains all the commands to run the application
- `migrations_embed.go` - Needed for extra setup for integration tests

### How To Run <a name = "how_to_run"></a>
1. Run `make setup-local` to setup application's dependencies
2. Run `make run` to run the application

### Integration Tests <a name = "integration_test"></a>
1. Run `make integration-test` to setup integration's test.