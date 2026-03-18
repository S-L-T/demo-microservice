FROM golang:1.16.3-alpine AS build-env

RUN apk update
RUN apk add bash protoc=3.13.0-r2 protobuf-dev=3.13.0-r2
WORKDIR /protoc_builder

COPY . .
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
