
GOPATH:=$(shell go env GOPATH)
MODIFY=Mgithub.com/micro/go-micro/api/proto/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
    
	protoc --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/api/api.proto
    

.PHONY: build
build: proto

	GOOS=linux GOARCH=amd64 go build -o api-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t api-service:latest
