#!/bin/sh

set -e

echo "Waiting for CentralStore DB..."
until mysqladmin ping -h centralstore-db -u"$CENTRALSTORE_DB_USER" -p"$CENTRALSTORE_DB_PASSWORD" --silent; do
  sleep 1
done

echo "Creating centralstore_db if it doesn't exist..."
mysql -h centralstore-db -u"$CENTRALSTORE_DB_USER" -p"$CENTRALSTORE_DB_PASSWORD" -e "CREATE DATABASE IF NOT EXISTS centralstore_db;"

echo "Running CentralStore migrations..."
mysql -h centralstore-db -u"$CENTRALSTORE_DB_USER" -p"$CENTRALSTORE_DB_PASSWORD" centralstore_db < central_store/data/migrations/1_init.up.sql

echo "✅ All migrations complete."
