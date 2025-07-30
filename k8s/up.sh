#!/bin/bash
set -e

source .env

kubectl create secret generic citystore-db-creds \
  --from-literal=CITYSTORE_DB_USER="${CITYSTORE_DB_USER}" \
  --from-literal=CITYSTORE_DB_PASSWORD="${CITYSTORE_DB_PASSWORD}"

kubectl create secret generic centralstore-db-creds \
  --from-literal=CENTRALSTORE_DB_USER="${CENTRALSTORE_DB_USER}" \
  --from-literal=CENTRALSTORE_DB_PASSWORD="${CENTRALSTORE_DB_PASSWORD}"

kubectl create secret generic citystore-db-password \
  --from-file=MYSQL_ROOT_PASSWORD_FILE=../passwords/citystore_db_password.txt

kubectl create secret generic centralstore-db-password \
  --from-file=MYSQL_ROOT_PASSWORD_FILE=../passwords/centralstore_db_password.txt

kubectl apply -f configs/city-config.yaml
kubectl apply -f configs/central-config.yaml
kubectl apply -f configs/city-db-config.yaml
kubectl apply -f configs/central-db-config.yaml
kubectl apply -f volumes.yaml
kubectl apply -f deployments.yaml
kubectl apply -f services.yaml
kubectl apply -f ingress.yaml

echo "Starting dashboard..."
minikube dashboard
