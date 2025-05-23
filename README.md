# SimpleBank 🏦

Backend de um sistema bancário simples desenvolvido em Go, utilizando PostgreSQL como banco de dados e `sqlc` para geração de código type-safe para acesso aos dados.

## 🌟 Funcionalidades

* **Contas:** Criar, ler, listar, atualizar e deletar contas.
* **Movimentações (Entradas):** Registrar entradas e saídas em contas.
* **Transferências:** Realizar transferências de valores entre contas.

## 🛠️ Tecnologias

* Go
* PostgreSQL
* `sqlc`
* `golang-migrate` (ou similar para migrações)
* Testes: `testing`, `testify/require`

## 🚀 Configuração Rápida

1.  **Pré-requisitos:**
    * Go (1.18+)
    * PostgreSQL
    * `sqlc` CLI ([sqlc.dev](https://sqlc.dev/))
    * `golang-migrate` CLI ([github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate))
    * Docker

2.  **Banco de Dados:**
    * Inicie o PostgreSQL.
    * Crie um banco de dados (ex: `simple_bank`) e um usuário/senha (ex: `root`/`root`).
    * A string de conexão padrão para testes é `postgresql://root:root@localhost:5432/simple_bank?sslmode=disable`.
    ```bash
    # Comando Para iniciar o PostgreSQL com docker 
    make postgres
    ```

3.  **Migrações:**
    Execute as migrações para criar o esquema do banco:
    ```bash
    make migrateup
    # Alternativa de não utilizar o Makefile: Substitua YOUR_DB_URL pela sua string de conexão completa
    migrate -path db/migration -database "YOUR_DB_URL" up
    ```
    Isso aplicará os scripts de `db/migration/` (ex: `000001_init_schema.up.sql`).

4.  **Gerar Código Go:**
    A partir da raiz do projeto, gere o código da camada de banco:
    ```bash
    make sqlc
    # Alternativa de não utilizar o Makefile:
    sqlc generate
    ```
    Isso usa o `sqlc.yaml` para converter as queries em `db/query/` para código Go em `db/sqlc/`.

## 🧪 Executando os Testes

Para rodar os testes da camada de banco de dados:

```bash
make test
# A partir da raiz do projeto (Alternativa)
go test ./db/sqlc/... -v