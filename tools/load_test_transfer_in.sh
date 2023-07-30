#!/bin/sh

k6 run ../loadtest/transfer_in.js --vus 100 --duration 1m