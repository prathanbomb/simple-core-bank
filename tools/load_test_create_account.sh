#!/bin/sh

bombardier -c 100 -n 1000000 -m POST http://localhost:8080/api/create-account -b '{"account_name": "พี่บอม เทพซ่า"}' --timeout=20s