#!/usr/bin/env bash

echo "restoring from backup"

PGPASSWORD="$POSTGRESQL_PASSWORD" \
  pg_restore \
  --clean \
  --if-exists \
  --no-owner \
  --no-privileges \
  --format c \
  --host localhost \
  --port 5432 \
  --username "$POSTGRESQL_USERNAME" \
  --dbname "$POSTGRESQL_DATABASE" \
  --verbose \
  "$SEED_FILENAME"

echo "backup restore complete"
