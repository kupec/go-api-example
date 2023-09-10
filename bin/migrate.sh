#!/bin/bash
set -ex

ENV_PATH="${ENV_PATH:-".env"}"
source "$ENV_PATH"

DATABASE="$DATABASE_DRIVER://$DATABASE_SOURCE"
migrate -path models/migrations -database "$DATABASE" -verbose "$@"
