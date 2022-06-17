
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go install github.com/micro/micro/v3/cmd/protoc-gen-micro@latest
	go install github.com/micro/micro/v3/cmd/protoc-gen-openapi@latest

.PHONY: api
api:
	protoc --openapi_out=. --proto_path=. proto/category.proto

.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=./proto/category/ --go_out=:./proto/category proto/category/category.proto
	
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o category

.PHONY: docker
docker:
	docker build . -t category:latest

.PHONY: dockerRun
dockerRun:
	docker run -d category
