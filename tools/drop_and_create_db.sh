#!/bin/sh

docker exec -i simple-core-bank-postgres psql -U postgres -c "drop database if exists simple_core_bank" &&
docker exec -i simple-core-bank-postgres psql -U postgres -c "create database simple_core_bank"
