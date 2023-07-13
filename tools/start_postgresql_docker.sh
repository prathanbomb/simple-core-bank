#!/bin/sh

docker run \
    -d \
    -p 5433:5432 \
    -e TZ=Asia/Bangkok \
    -e POSTGRES_PASSWORD=postgres \
    --name simple-core-bank-postgres postgres:14-alpine
