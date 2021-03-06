#!/bin/bash

echo "Will set environment variables (Using Linux Command Export):"

export SERVER_ADDRESS=localhost \
SERVER_PORT=8001 \
DB_USER=root \
DB_PASSWD=codecamp \
DB_ADDR=localhost \
DB_PORT=3306 \
DB_NAME=banking \
AUTH_MIDDLEWARE_SERVER_ADDRESS=localhost \
AUTH_MIDDLEWARE_SERVER_PORT=8181

echo "Will start Banking App"

cd ./banking
go run main.go
