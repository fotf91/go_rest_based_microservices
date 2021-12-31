#!/bin/bash

echo "Will set environment variables (Using Linux Command Export):"

export SERVER_ADDRESS=localhost \
SERVER_PORT=8181 \
DB_USER=root \
DB_PASSWD=codecamp \
DB_ADDR=localhost \
DB_PORT=3306 \
DB_NAME=banking \

echo "Will start Middleware Auth Server"

cd ./banking-auth
go run main.go
