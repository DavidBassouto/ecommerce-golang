# ecommerce-golang

API de e-commerce escrita em Go utilizando Gin, GORM (PostgreSQL), JWT e integração S3 (via LocalStack).

**Visão Geral**

- Fornece um servidor HTTP básico com rota de saúde (`GET /health`).
- Camadas de serviço e utilitários já implementadas para autenticação, domínio de produtos, carrinho e pedidos.
- Estrutura principal:
  - `cmd/api` (entrypoint)
  - `internal/config`, `internal/database`, `internal/server`, `internal/services`, `internal/models`, `internal/dto`, `internal/utils`
  - `docker/docker-compose.yml` para serviços locais (Postgres e LocalStack)
  - `db/migrations` para migrações SQL

**Requisitos**

- Go 1.25+ instalado.
- Docker e Docker Compose.
- Make (opcional). Se não usar Make, os comandos podem ser executados diretamente com `go` e `docker compose`.
- Ferramenta de migração `migrate` (opcional) instalada localmente.

**Instalação**

- Clone o repositório.
- Copie o arquivo de exemplo de ambiente: `cp .env.exemple .env` (ou crie manualmente).
- Preencha as variáveis de ambiente no `.env` (sem commitar segredos).

**Variáveis de Ambiente (nomes esperados)**

- Servidor: `SERVER_PORT`, `GIN_MODE`
- Banco: `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_SSL_MODE`
- JWT: `JWT_SECRET_KEY`, `JWT_EXPIRES_IN`, `JWT_REFRESH_EXPIRES_IN`
- AWS/S3: `AWS_REGION`, `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_S3_BUCKET_NAME`, `AWS_S3_ENDPOINT`
- Uploads: `UPLOAD_PATH`, `UPLOAD_MAX_FILE_SIZE`
- Observação: os nomes devem seguir o que o código espera. Não inclua valores reais no README; defina-os apenas no seu `.env`.

**Serviços Locais (Docker)**

- Subir serviços:
  - Com Make: `make docker-up`
  - Sem Make: `docker compose -f docker/docker-compose.yml up -d`
- Derrubar serviços:
  - Com Make: `make docker-down`
  - Sem Make: `docker compose -f docker/docker-compose.yml down`
- Dica: portas podem variar conforme seu ambiente; ajuste mapeamentos em `docker/docker-compose.yml` se necessário.

**Migrações de Banco**

- Com Make:
  - Up: `make migrate-up`
  - Down: `make migrate-down`
- Sem Make: utilize sua própria string de conexão com a CLI `migrate` apontando para `db/migrations`.
- Observação: ajuste credenciais e host para refletir seu `.env`/container local. Evite strings de conexão fixas fora do seu ambiente.

**Executando a Aplicação**

- Com Make:
  - Dev: `make dev`
  - Run: `make run`
  - Build: `make build`
- Sem Make:
  - `go run ./cmd/api`
  - Binário: `go build -o bin/app ./cmd/api`
- A porta do servidor é definida por `SERVER_PORT`. A rota de saúde está em `/health`.

**Qualidade de Código**

- Formatar: `make format` (ou `gofmt`/`goimports` diretamente).
- Lint: `make lint` (golangci-lint).

**Notas de Portabilidade**

- Todos os caminhos devem ser relativos à raiz do projeto; evite caminhos absolutos específicos do seu computador.
- Em Windows, macOS e Linux, os comandos são equivalentes. Caso não tenha `make`, execute os comandos correspondentes com `go` e `docker compose`.
- Não commite `.env` ou segredos. Use `.env.exemple` apenas como referência de chaves.

**Rotas Atuais**

- `GET /health` — retorna `{ "message": "OK" }`.

**Próximos Passos**

- Implementar handlers/rotas para autenticação, produtos, carrinho e pedidos utilizando as camadas de serviço e utilitários já existentes.
