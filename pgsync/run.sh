#!/usr/bin/env bash

set -e

./wait-for-it.sh "$PG_HOST:$PG_PORT" -t 60

./wait-for-it.sh "$ELASTICSEARCH_HOST:$ELASTICSEARCH_PORT" -t 60

./wait-for-it.sh "$REDIS_HOST:$REDIS_PORT" -t 60

"$HOME/.local/bin/bootstrap" --config "$SCHEMA_CONFIG"

"$HOME/.local/bin/pgsync" --config "$SCHEMA_CONFIG" --daemon
