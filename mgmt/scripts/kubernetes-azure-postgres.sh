minikube start

kubectl apply -f ./kubernetes-minikube-azure-postgres/mgmt.yaml

sleep 60

minikube service mgmt-frontend-service