#!/bin/bash

set -e

echo "run db migration"
source /app/app.env
echo "db source : $DB_SOURCE"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"