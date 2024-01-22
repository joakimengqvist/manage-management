#!/bin/bash

# List of your microservices image names
microservices=(
    "broker"
    "authentication"
    "project"
    "external-company"
    "economics"
    "notes"
    "invoice"
    "product"
    )

# Your Docker Hub username
username="joakimengqvist"

# Loop through each microservice and push the image to Docker Hub
for service in "${microservices[@]}"; do
    docker tag "mgmt-$service-service" "joakimengqvist/mgmt-$service-service"
    docker push "joakimengqvist/mgmt-$service-service"
done

docker tag "mgmt-frontend" "joakimengqvist/mgmt-frontend"
docker push "joakimengqvist/mgmt-frontend"