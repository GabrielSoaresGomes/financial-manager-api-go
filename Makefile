.DEFAULT_GOAL := help

help:
	@echo "Comandos disponiveis:"
	@echo "  make setup     - Configura o ambiente local para rodar o projeto"
	@echo "  make dev       - Sobe a aplicacao com hot reload (Air)"
	@echo "  make run       - Sobe a aplicacao (migrations rodam automaticamente)"
	@echo "  make build     - Gera o binario em bin/"
	@echo "  make db/up     - Sobe o banco de dados via Docker"
	@echo "  make db/down   - Derruba o banco de dados"
	@echo "  make db/logs   - Exibe os logs do Postgres"
	@echo "  make tidy      - Limpa e sincroniza as dependencias"
	@echo "  make test/http - Roda os testes HTTP via Docker"

setup:
	@if [ ! -f .env.local ]; then \
		printf 'APP_ENV=local\nAPP_PORT=2004\n\nDB_HOST=localhost\nDB_PORT=5432\nDB_USER=postgres\nDB_PASSWORD=12345678\nDB_NAME=financial_manager\n' > .env.local; \
		echo ".env.local criado com valores padrao"; \
	else \
		echo ".env.local ja existe, pulando..."; \
	fi
	go mod download
	go install github.com/air-verse/air@latest
	docker compose up -d
	@echo "Setup concluido! Rode 'make dev' para iniciar."

dev:
	air

run:
	go run cmd/main.go

build:
	go build -o bin/app cmd/main.go

db/up:
	docker compose up -d

db/down:
	docker compose down

db/logs:
	docker compose logs -f postgres

tidy:
	go mod tidy

test/http:
	docker run --rm -v ${CURDIR}:/workdir jetbrains/intellij-http-client --env-file http/http-client.env.json --env local $(wildcard http/*.http) -D
