#!/bin/bash

echo "Will set environment variables (Using Linux Command Export):"

export SERVER_ADDRESS=localhost \
SERVER_PORT=8001 \
DB_USER=root \
DB_PASSWD=codecamp \
DB_ADDR=localhost \
DB_PORT=3306 \
DB_NAME=banking \

echo "Will Run main.go File"

cd ./banking
go run main.go
