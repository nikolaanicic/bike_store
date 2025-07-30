#!/bin/sh
set -e

sleep 20
echo "Creating citystore_db if it doesn't exist..."
mysql -h citystore-db -u"$CITYSTORE_DB_USER" -p"$CITYSTORE_DB_PASSWORD" --ssl=0 -e "CREATE DATABASE IF NOT EXISTS citystore_db;"

echo "Running CityStore migrations..."
mysql -h citystore-db -u"$CITYSTORE_DB_USER" -p"$CITYSTORE_DB_PASSWORD" --ssl=0 citystore_db < city_store/data/migrations/1_init.up.sql

echo "✅ All migrations complete. exiting..."

echo "Starting the backend server"
./citystore_backend
