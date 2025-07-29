kubectl create secret generic citystore-db-password \
  --from-file=MYSQL_ROOT_PASSWORD=../passwords/citystore_db_password.txt

kubectl create secret generic centralstore-db-password \
  --from-file=MYSQL_ROOT_PASSWORD=../passwords/centralstore_db_password.txt

kubectl apply -f configs/city-config.yaml
kubectl apply -f configs/central-config.yaml
kubectl apply -f configs/city-db-config.yaml
kubectl apply -f configs/central-db-config.yaml
kubectl apply -f volumes.yaml
kubectl apply -f deployments.yaml
kubectl apply -f services.yaml
sleep 1
kubectl apply -f ingress.yaml

# clear
echo "Starting"
minikube dashboard