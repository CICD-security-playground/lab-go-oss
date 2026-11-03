APP_NAME := lab-go-oss
BINARY_DIR := bin
BINARY := $(BINARY_DIR)/$(APP_NAME)
DOCKER_IMAGE ?= $(APP_NAME):local

.PHONY: build test run docker-build security-stubs clean

build:
	mkdir -p $(BINARY_DIR)
	go build -trimpath -o $(BINARY) ./cmd/server

test:
	go test ./...

run:
	go run ./cmd/server

docker-build:
	docker build -t $(DOCKER_IMAGE) .

security-stubs:
	./scripts/sast-stub.sh
	./scripts/sca-stub.sh
	./scripts/secrets-stub.sh

clean:
	rm -rf $(BINARY_DIR) coverage.out
