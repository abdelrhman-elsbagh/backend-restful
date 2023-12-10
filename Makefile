# Makefile

DOCKERFILE = Dockerfile
DOCKER_COMPOSE_FILE = docker-compose.yml

# Extract the image name from the Dockerfile
IMAGE_NAME = $(shell grep -m 1 "^FROM" $(DOCKERFILE) | cut -d ' ' -f 2)

.PHONY: build run clean

build:
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) .

run:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up

clean:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down --volumes --remove-orphans
	docker rmi $(IMAGE_NAME) || true

# Target to build and run the Docker container
start: build run

# Target to build, run, and then clean up
start-clean: build run clean
