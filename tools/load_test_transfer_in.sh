#!/bin/sh

bombardier -c 100 -n 10000 -m POST http://localhost:8080/api/transfer-in-for-load-test -b '{"max_account_no": 1000000, "amount": 1000000}' --timeout=20s