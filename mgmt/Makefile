BROKER_BINARY=brokerApp
AUTH_BINARY=authApp
PROJECT_BINARY=projectApp
LOGGER_BINARY=loggerApp
LISTENER_BINARY=listenerApp
NOTES_BINARY=notesApp
ECONOMICS_BINARY=economicsApp
EXTERNAL_COMPANY_BINARY=externalCompanyApp
PRODUCT_BINARY=productApp
INVOICE_BINARY=invoiceApp


## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

go_build: build_broker build_auth build_project build_notes build_economics build_external_company build_product build_invoice

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker build_auth build_project build_notes build_economics build_external_company build_product build_invoice build_front
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose compose-mgmt-local-postgres.yml up --build -d
	@echo "Docker images built and started!"
 
## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker compose down
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

## build_auth: builds the broker binary as a linux executable
build_auth:
	@echo "Building auth binary..."
	cd ../authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Done!"

## build_project: builds the project binary as a linux executable
build_project:
	@echo "Building project binary..."
	cd ../project-service && env GOOS=linux CGO_ENABLED=0 go build -o ${PROJECT_BINARY} ./cmd/api
	@echo "Done!"

## build_notes: builds the notes binary as a linux executable
build_notes:
	@echo "Building notes binary..."
	cd ../notes-service && env GOOS=linux CGO_ENABLED=0 go build -o ${NOTES_BINARY} ./cmd/api
	@echo "Done!"

## build_economics: builds the economics binary as a linux executable
build_economics:
	@echo "Building economics binary..."
	cd ../economics-service && env GOOS=linux CGO_ENABLED=0 go build -o ${ECONOMICS_BINARY} ./cmd/api
	@echo "Done!"

## build_external_company: builds the external company binary as a linux executable
build_external_company:
	@echo "Building external_company binary..."
	cd ../external-company-service && env GOOS=linux CGO_ENABLED=0 go build -o ${EXTERNAL_COMPANY_BINARY} ./cmd/api
	@echo "Done!"

## build_external_company: builds the external company binary as a linux executable
build_product:
	@echo "Building product binary..."
	cd ../product-service && env GOOS=linux CGO_ENABLED=0 go build -o ${PRODUCT_BINARY} ./cmd/api
	@echo "Done!"

## build_invoice: builds the invoice binary as a linux executable
build_invoice:
	@echo "Building invoice binary..."
	cd ../invoice-service && env GOOS=linux CGO_ENABLED=0 go build -o ${INVOICE_BINARY} ./cmd/api
	@echo "Done!"

## build_logger: builds the logger binary as a linux executable
build_logger:
	@echo "Building logger binary..."
	cd ../logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./cmd/api
	@echo "Done!"

## build_logger: builds the listener binary as a linux executable
build_listener:
	@echo "Building listener binary..."
	cd ../listener-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LISTENER_BINARY} .
	@echo "Done!"

## start: starts the front end
start_front_dev:
	@echo "Starting front end development"
	cd ../front-end && yarn dev

## start: starts the front end
start_front_prod:
	@echo "Starting front end production build"
	cd ../front-end && yarn build

## stop: stop the front end
stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped front end!"