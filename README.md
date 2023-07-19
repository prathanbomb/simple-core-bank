# simple-core-bank

## Prerequisites

- Make
- Docker 24 or later
- Go 1.19 or later
- PostgreSQL 14 or later

## Getting started

1. Start PostgreSQL docker

```sh
make start-db
```

2. Serve HTTP API

```sh
make run
```

## Load test

1. Load test create account API

```sh
./tools/load_test_create_account.sh 
```