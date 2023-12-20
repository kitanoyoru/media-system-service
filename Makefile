GOPATH := $(shell go env GOPATH)

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o lab ./cmd/server/main.go

.PHONY: run-local
run-local:
	@source ./.env && make build && ./lab
