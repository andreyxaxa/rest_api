.PHONY: build
build:
		go build -v ./cmd/apiserver

.PHONY: test
test:
		go test -v ./...

.PHONY: docker.run
docker.run:
		docker compose up -d

.PHONY: docker.rebuild
docker.rebuild:
		docker compose up -d --build

.PHONY: docker.down
docker.down:
		docker compose down

.DEFAULT_GOAL := build