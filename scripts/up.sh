#!/bin/bash

set -e

source .env

echo "Bringing up database containers..."
docker compose up --build -d --force-recreate centralstore_db
docker compose up --build -d --force-recreate citystore_db

sleep 15  # Wait for the databases to be ready

echo "Creating databases if they do not exist..."
docker exec citystore_db mysql -u"$CITYSTORE_DB_USER" -p"$CITYSTORE_DB_PASSWORD" -e "CREATE DATABASE IF NOT EXISTS citystore_db;"
docker exec centralstore_db mysql -u"$CENTRALSTORE_DB_USER" -p"$CENTRALSTORE_DB_PASSWORD" -e "CREATE DATABASE IF NOT EXISTS centralstore_db;"

echo "Running migrations..."
docker exec -i citystore_db mysql -u"$CITYSTORE_DB_USER" -p"$CITYSTORE_DB_PASSWORD" citystore_db < ../city_store/data/migrations/1_init.up.sql 
docker exec -i centralstore_db mysql -u"$CENTRALSTORE_DB_USER" -p"$CENTRALSTORE_DB_PASSWORD" centralstore_db < ../central_store/data/migrations/1_init.up.sql


echo "Starting services..."
docker compose --verbose up -d --build centralstore
docker compose --verbose up -d --build citystore

echo "✅ All services are up and running."