#!/bin/sh

set -e

echo "Waiting for CityStore DB..."
until mysqladmin ping -h localhost -u"$CITYSTORE_DB_USER" -p"$CITYSTORE_DB_PASSWORD" --silent; do
  sleep 1
done

echo "Creating citystore_db if it doesn't exist..."
mysql -h citystore-db -u"$CITYSTORE_DB_USER" -p"$CITYSTORE_DB_PASSWORD" -e "CREATE DATABASE IF NOT EXISTS citystore_db;"

echo "Running CityStore migrations..."
mysql -h citystore-db -u"$CITYSTORE_DB_USER" -p"$CITYSTORE_DB_PASSWORD" citystore_db < city_store/data/migrations/1_init.up.sql

echo "✅ All migrations complete."
