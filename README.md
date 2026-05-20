# Financial Manager API

API REST para gerenciamento financeiro, desenvolvida em Go.

## Requisitos

- [Go 1.25+](https://golang.org/dl/)
- [Docker](https://www.docker.com/) e Docker Compose
- [Make](https://community.chocolatey.org/packages/make) (no Windows: `choco install make` como administrador)
- [Git](https://git-scm.com/)

## Configuração

1. Clone o repositório e entre na pasta:
   ```bash
   git clone <repo-url>
   cd financial-manager-api-go
   ```

2. Copie o arquivo de variáveis de ambiente e preencha os valores:
   ```bash
   cp .env.example .env.local
   ```

3. Baixe as dependências:
   ```bash
   make tidy
   ```

## Como rodar

1. Suba o banco de dados:
   ```bash
   make db/up
   ```

2. Suba a aplicação (as migrations rodam automaticamente na inicialização):
   ```bash
   make run
   ```

   Ou com hot reload (requer [Air](https://github.com/air-verse/air)):
   ```bash
   go install github.com/air-verse/air@latest
   make dev
   ```

A API estará disponível em `http://localhost:2004`.

## Testes HTTP

Com a aplicação rodando, execute os testes HTTP (requer Docker):

```bash
make test/http
```

Os testes ficam na pasta `http/` e utilizam o ambiente `local` definido em `http/http-client.env.json`.

## Comandos disponíveis

```
make dev       - Sobe a aplicação com hot reload (Air)
make run       - Sobe a aplicação
make build     - Gera o binário em bin/
make db/up     - Sobe o banco de dados
make db/down   - Derruba o banco de dados
make db/logs   - Exibe os logs do Postgres
make tidy      - Limpa e sincroniza as dependências
make test/http - Roda os testes HTTP via Docker
```
