# Simple Bank

- A simple bank backend application built on Go.

# Features

- Create and manage accounts: Owner, balance and currency
- Record all balance changes (deposits and withdrawals): create an account entry for each change.
- Money transfer transaction: Perform transfer between account consistently within a transaction.

# Development Process.

-Design database using dbdiagram.io application and export the design to both
sql and png format.

- Using golang-migrate as migration too. This enable developers to easily modify the schemas
  commands:

```sh
  $ brew install golang-migrate
```

- Golang-migrate Essential commands:

  - create
  - up
  - down
  - goto

- Command for creating migration

```sh
   $ brew install golang-migrate
   $ migrate create -ext sql -dir db/migration -seq init
   $ migrate -path db/migration -database 'postgresql://admin:y7jHf&DNWG15@localhost:5030/main?sslmode=disable' -verbose up
```

# Enable SSL on the Server

If you want to use SSL (for example, in production environments), you need to enable SSL on the PostgreSQL server and configure it properly. To enable SSL, you’ll need to:

Modify your postgresql.conf file:
Set ssl = on
Ensure ssl_cert_file and ssl_key_file are set to valid certificate and key file paths.
Restart the PostgreSQL service to apply the changes.
Example of the relevant settings in postgresql.conf:

ini
Copy code
ssl = on
ssl_cert_file = '/path/to/server.crt'
ssl_key_file = '/path/to/server.key'

# CRUD

- CREATE
- READ
- UPDATE
- DELETE

Making Decision on what database to choose.

### Database/SQL (Go standard library)

- Very fast and straightforward
- Manual mapping SQL fields to variables
- Easy to make mistake, not caught until runtime. But you have test cases though?

### GORM

- CRUD functions already implemented, very short production code
- Must learn how to write queries using GORM functions.
- Run very slowly on high load:

### SQLX

- Quite fast and easy to use
- Field mapping via query text & struct tags
- Failure won't occur until runtime.

### SQLC

- Very fast and easy to use
- Write SQL code and automatically generate CRUD code.
- Catch SQL query errors before generating codes.
- Currently only support Postgres. MYSQL is experimental.

```sh
  brew install sqlc
  # commands: compile, generate, init
```

## Writing Unit Tests.

- By convention, test files are put inside the code or packages.
- Write unit tests for CRUD functions.
- Database driver for postgres: libpq
- Tools: testing, and testify
