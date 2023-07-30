#!/bin/sh

k6 run ../loadtest/transfer.js --vus 100 --duration 1m