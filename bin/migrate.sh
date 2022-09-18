#!/bin/sh
set -ex

ENV_PATH="${ENV_PATH:-".env"}"
source "$ENV_PATH"

migrate -path models/migrations -database "$DATABASE_URL" -verbose up
