# syntax=docker/dockerfile:1
FROM golang:1.18-alpine

RUN apk add --no-cache protoc

WORKDIR /api

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

ENV export PATH="$PATH:$(go env GOPATH)/bin"

COPY . ./
RUN pwd

RUN go mod tidy
RUN protoc -Iproto --go_out=. --go-grpc_out=. proto/shelter.proto

RUN go build -o /bin/app ./cmd/api/main.go

CMD ["/bin/app"]

