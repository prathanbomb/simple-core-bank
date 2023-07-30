#!/bin/sh

k6 run ../loadtest/create_account.js --vus 100 --duration 1m