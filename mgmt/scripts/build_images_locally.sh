#!/bin/bash

# List of your microservices image names
microserviceNames=(
    "broker"
    "authentication"
    "project"
    "external-company"
    "economics"
    "notes"
    "invoice"
    "product"
    )

for service in "${microserviceNames[@]}"; do
    docker build --no-cache -t "mgmt-$service-service" -f "./../$service-service/$service-service.dockerfile" "./../$service-service"
done

docker build --no-cache -t "mgmt-frontend" -f "./../front-end/dockerfile" "./../front-end"




