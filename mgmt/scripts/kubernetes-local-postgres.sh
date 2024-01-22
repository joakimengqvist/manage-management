docker compose -f compose-postgres-local.yml up -d

minikube start

kubectl apply -f ./kubernetes-minikube-local-postgres/mgmt.yaml

sleep 60

minikube service mgmt-frontend-service