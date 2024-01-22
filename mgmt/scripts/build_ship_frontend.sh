

docker build --no-cache -t "mgmt-frontend" -f "./../front-end/dockerfile" "./../front-end"


docker tag "mgmt-frontend" "joakimengqvist/mgmt-frontend"
docker push "joakimengqvist/mgmt-frontend"

