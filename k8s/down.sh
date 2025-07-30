# Desc: Stops and deletes all kubernetes resources
kubectl delete ingress --all
kubectl delete service --all
kubectl delete deployment --all
kubectl delete pvc --all
kubectl delete pv --all
kubectl delete secret --all
kubectl delete job --all
clear
echo "Stopped and deleted all kubernetes resources"