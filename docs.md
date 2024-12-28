# Simple Bank

- A simple bank backend application built on Go.

# Features

- Create and manage accounts: Owner, balance and currency
- Record all balance changes (deposits and withdrawals): create an account entry for each change.
- Money transfer transaction: Perform transfer between account consistently within a transaction.

# Development Process.

- Design database using dbdiagram.io application and export the design to both
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

## Golang DB Transactions

- Database Transactions: A single unit of work often made up of multiple db operations.
  Example: Transfer 10 USD from bank account1 to bank account 2

  - Create a transfer record with amount = 10
  - Create an account entry for account1 with amount = -10
  - Create an account entry for account2 with amount = +10
  - Substract 10 from the balance of account1
  - Add 10 to the balance of account2

- Why do we need db transaction?

  - to provide a reliable and consistent unit of work.
    even in the case of system failure.
  - to provide isolation between programs that access the database
    concurrently

- Database Transaction must satisfy ACID properties.

  - Atomicity: Either all operations complete successfully or the transactions
    fails and the db is unchanged.
  - Consistency: The db state must be valid after the transaction. All constraints must be satisfied.
  - Isolation: Concurrent transactions must not affect each other.
  - Durability: Data written by a successfully transaction must be recorded in persistent storage.

- Transactions Problems:
  - Deadlocks: occurs when two concurrent transactions cannot make progress because each one
    waits for the other to release a lock.
  - Solved problems with transaction.

## Implementing Continuous Integration (CI)

- Tools: Github Actions
  - Workflow: An automated procedure that contains one or more jobs.
    - Worflow is made of up one or more jobs.
    - A workflow is triggered by events happening on a repository or manually or a schedule.
    - Add .yml file for configuration.
  - Runner: This is a server that run the jobs.
    - Each job has its own runner.
    - You can custom or Github runners.
  - Job: A set of steps which execute on the same runner.
    - Normally jobs run in parallel
    - Dependended jobs run serially.
  - Step: An individual tasks.
    - An job can have one or or more steps.
    - A step contains one action.
    - Action is a standalone command. A step can have one or more actions.

## Developing Restful API using gin-gonic web framework.

- Provides:Routing, Middleware and Authentication features.

```sh
   # Install gin-gonic
   go get -u github.com/gin-gonic/gin
```

## Enviroment Variables and Configurations:

Tools: Viper

- Viper is the most popular tool for reading environment variables
  and configurations.
- Viper has the following features:
  - Find, load, unmarshal config files.
  - Read config from environment variables or flags.
  - Read config from remote system.
  - Live watching and writing config file.

# Hashing Password for Security

- Tools: Bcrypt
- Bcrypt package in Go implement provos and Mazieres bcrypt adaptive hashing algorithm. (http://www.usenix.org/event/usenix99/provos/provos.pdf)

# Authention and Authorization:

Tools: JWT token package and Paseto.
