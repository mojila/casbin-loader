#!/bin/sh

export CASBIN_HOST=localhost:50051
export DB_HOST=postgresql
export DB_PORT=5432
export DB_USERNAME=admin
export DB_PASSWORD=admin
export DB_NAME=app

./casbin-loader