# Clean architecture: Go

A Clean Architecture template for a Rest API in Go

# How it works

Motivations, explanations, requirements & more details in my article [Practical Clean Architecture in Typescript, Rust & Python](https://dev.to/msc29/practical-clean-architecture-in-typescript-rust-python-3a6d)

# Installing

```bash
cargo build
```

# Database setup

It's currently configured to run with PostgreSQL through Gorm (ORM), but this being clean architecture feel free to change it :)

I suggest

- PostgreSQL [in docker](https://hub.docker.com/_/postgres/)
- pgAdmin [install](https://www.pgadmin.org/download/pgadmin-4-apt/)

# Running

TODO

define the environment on which we're running by adding `ENV=<env>`, which will use the `.env.<env>` file

```bash
ENV=dev go run ./src/main.go
```

# Code quality & security

Used in CI/CD

```bash
go fmt <go file or package>
```

# Testing

run the mock api (https://github.com/tkc/go-json-server)

```bash
cd tests/integration_tests/mock_api && go-json-server
```

run the tests

```bash
ENV=test go test ./...
```

# API Documentation

TODO
