BROKER_BINARY=brokerApp
AUTH_BINARY=authApp
PROJECT_BINARY=projectApp
LOGGER_BINARY=loggerApp
MAIL_BINARY=mailerApp
LISTENER_BINARY=listenerApp
NOTES_BINARY=notesApp
ECONOMICS_BINARY=economicsApp


## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker build_auth build_project build_notes build_economics build_logger build_mail build_listener
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
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

## build_logger: builds the broker binary as a linux executable
build_logger:
	@echo "Building logger binary..."
	cd ../logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./cmd/api
	@echo "Done!"

## build_mail: builds the broker binary as a linux executable
build_mail:
	@echo "Building mail binary..."
	cd ../mail-service && env GOOS=linux CGO_ENABLED=0 go build -o ${MAIL_BINARY} ./cmd/api
	@echo "Done!"

## build_logger: builds the broker binary as a linux executable
build_listener:
	@echo "Building listener binary..."
	cd ../listener-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LISTENER_BINARY} .
	@echo "Done!"

## build_front: builds the frone end binary
build_front:
	@echo "Building front end binary..."
	cd ../front-end && yarn dev
	@echo "Done!"

## start: starts the front end
start: build_front
	@echo "Starting front end"
	cd ../front-end && yarn dev

## stop: stop the front end
stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped front end!"