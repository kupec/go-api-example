#!/bin/sh
set -ex

ENV_PATH="${ENV_PATH:-".env"}"
source "$ENV_PATH"

migrate create -ext sql -dir models/migrations -seq init_schema
